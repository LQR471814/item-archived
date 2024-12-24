package service

import (
	"errors"
	"fmt"
	v1 "item-archived/api/v1"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

type imageExt struct {
	ext    string
	format v1.ImageFormat
}

var image_extensions = []imageExt{
	{"jpg", v1.ImageFormat_JPG},
	{"png", v1.ImageFormat_PNG},
	{"gif", v1.ImageFormat_GIF},
	{"svg", v1.ImageFormat_SVG},
}

func parseFilename(filename string) (id string, tags []string, isContainer bool, err error) {
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		return "", nil, false, fmt.Errorf("parse filename: invalid filename \"%s\"", filename)
	}
	id = parts[0]
	isContainer = parts[len(parts)-1] == "container"
	for i := 1; i < len(parts)-1; i++ {
		tags = append(tags, parts[i])
	}
	return id, tags, isContainer, nil
}

func formatFilename(id string, tags []string, isContainer bool) string {
	segments := append([]string{id}, tags...)
	if isContainer {
		segments = append(segments, "container")
	} else {
		segments = append(segments, "item")
	}
	return strings.Join(segments, ".")
}

func readEntryMeta(fpath string) (*v1.EntryMetadata, error) {
	_, filename := filepath.Split(fpath)
	id, tags, _, err := parseFilename(filename)
	if err != nil {
		return nil, fmt.Errorf("readEntryMeta: %w", err)
	}

	var imgFormat *v1.ImageFormat
	var img []byte
	for _, ext := range image_extensions {
		imagePath := filepath.Join(fpath, fmt.Sprintf("image.%s", ext.ext))

		img, err = os.ReadFile(imagePath)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			slog.Warn("failed to read image file", "filepath", imagePath, "err", err)
			continue
		}
		imgFormat = &ext.format
	}

	var description *string
	descPath := filepath.Join(fpath, "description.txt")
	descContents, err := os.ReadFile(descPath)
	if err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			slog.Warn("failed to read description", "filepath", descPath, "err", err)
		}
	} else {
		descContentsStr := string(descContents)
		description = &descContentsStr
	}

	return &v1.EntryMetadata{
		Id:          id,
		Tags:        tags,
		Description: description,
		Image:       img,
		ImageFormat: imgFormat,
	}, nil
}

func writeEntryMeta(fpath string, meta *v1.EntryMetadata, isContainer bool) error {
	filename := formatFilename(meta.GetId(), meta.GetTags(), isContainer)

	err := os.Mkdir(filepath.Join(fpath, filename), 0777)
	if err != nil {
		return fmt.Errorf("writeEntryMeta: %w", err)
	}

	err = os.WriteFile(
		filepath.Join(fpath, filename, "description.txt"),
		[]byte(meta.GetDescription()),
		0600,
	)
	if err != nil {
		return fmt.Errorf("writeEntryMeta: %w", err)
	}

	if len(meta.GetImage()) > 0 {
		imgExt := ""
		for _, ext := range image_extensions {
			if ext.format == meta.GetImageFormat() {
				imgExt = ext.ext
				break
			}
		}
		err = os.WriteFile(
			filepath.Join(fpath, filename, fmt.Sprintf("image.%s", imgExt)),
			meta.GetImage(),
			0600,
		)
		if err != nil {
			return fmt.Errorf("writeEntryMeta: %w", err)
		}
	}

	return nil
}

func readChildren(fpath string) (items []string, containers []string, err error) {
	entries, err := os.ReadDir(fpath)
	if err != nil {
		return nil, nil, fmt.Errorf("readChildren: %w", err)
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		if strings.HasSuffix(e.Name(), ".container") {
			containers = append(containers, e.Name())
			continue
		}
		if strings.HasSuffix(e.Name(), ".item") {
			items = append(items, e.Name())
			continue
		}
	}
	return items, containers, nil
}
