package service

import (
	"context"
	"fmt"
	v1 "item-archived/api/v1"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"connectrpc.com/connect"
)

/*

All directories/files should be lowercase.

A directory ending in `.item` represents a single item.

- `image.{jpg,png,gif,svg}` - an image of the item
- `description.txt` - a description of the item

A directory ending in `.container` represents a container that contains multiple items.

- `*.item` - items inside this container
- `*.container` - other containers inside this container
- `image.{jpg,png,gif,svg}` - an image of the container
- `description.txt` - a description of the container

You can add tags to a item or container by adding more extensions to the filename like this: `some_cool_thing.multiple.fruit.item`.

In this case, the `some_cool_thing` item has tags `multiple` and `fruit` applied to it.

*/

type Service struct {
	dir string
}

func NewService(dir string) (Service, error) {
	_, dirname := filepath.Split(dir)
	if !strings.HasSuffix(dirname, ".container") {
		return Service{}, fmt.Errorf(
			"NewService: you must use a directory with a .container extension as the root, '%s' does not have such an extension.",
			dirname,
		)
	}
	return Service{dir: dir}, nil
}

func (s Service) Read(ctx context.Context, req *connect.Request[v1.ReadRequest]) (*connect.Response[v1.ReadResponse], error) {
	path := req.Msg.GetPath()
	fpath := filepath.Join(append([]string{s.dir}, path...)...)

	slog.Debug("reading dir", "dir", fpath)

	_, err := os.Stat(fpath)
	if err != nil {
		return nil, err
	}
	meta, err := readEntryMeta(fpath)
	if err != nil {
		return nil, err
	}

	var children *v1.ReadResponse_Children
	if len(path) == 0 || strings.HasSuffix(path[len(path)-1], ".container") {
		items, containers, err := readChildren(fpath)
		if err != nil {
			return nil, err
		}
		children = &v1.ReadResponse_Children{
			ItemNames:      items,
			ContainerNames: containers,
		}
	}

	return &connect.Response[v1.ReadResponse]{
		Msg: &v1.ReadResponse{
			Metadata: meta,
			Children: children,
		},
	}, nil
}

func (s Service) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	path := req.Msg.GetPath()
	meta := req.Msg.GetMetadata()
	createContainer := req.Msg.GetCreateContainer()

	err := writeEntryMeta(filepath.Join(path...), meta, createContainer)
	if err != nil {
		return nil, err
	}

	return &connect.Response[v1.CreateResponse]{
		Msg: &v1.CreateResponse{},
	}, nil
}

func (s Service) Move(ctx context.Context, req *connect.Request[v1.MoveRequest]) (*connect.Response[v1.MoveResponse], error) {
	src := filepath.Join(req.Msg.GetSrc()...)
	dst := filepath.Join(req.Msg.GetDest()...)
	err := os.Rename(src, dst)
	if err != nil {
		return nil, err
	}
	return &connect.Response[v1.MoveResponse]{
		Msg: &v1.MoveResponse{},
	}, nil
}

func (s Service) Delete(ctx context.Context, req *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	path := filepath.Join(req.Msg.GetPath()...)
	err := os.RemoveAll(path)
	if err != nil {
		return nil, err
	}
	return &connect.Response[v1.DeleteResponse]{
		Msg: &v1.DeleteResponse{},
	}, nil
}

func (s Service) Search(ctx context.Context, req *connect.Request[v1.SearchRequest]) (*connect.Response[v1.SearchResponse], error) {

}
