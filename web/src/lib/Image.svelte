<script lang="ts">
  import { ImageFormat } from "../api/v1/api_pb";

  let {
    className,
    alt,
    bytes,
    format,
  }: {
    className: string;
    alt: string;
    bytes: ArrayBuffer;
    format: ImageFormat;
  } = $props();

  let old: string | undefined;
  const src = $derived.by(() => {
    if (old) {
      URL.revokeObjectURL(old);
    }

    let mime: string;
    switch (format) {
      case ImageFormat.JPG:
        mime = "image/jpeg";
        break;
      case ImageFormat.GIF:
        mime = "image/gif";
        break;
      case ImageFormat.PNG:
        mime = "image/png";
        break;
      case ImageFormat.SVG:
        mime = "image/svg+xml";
        break;
    }

    return URL.createObjectURL(new Blob([bytes], { type: mime }));
  });
</script>

<img class={className} {src} {alt} />
