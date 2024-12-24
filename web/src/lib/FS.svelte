<script lang="ts">
import type { Archive } from "../archive"

let { archive }: {
  archive: Archive
} = $props()

let cwd = $state<string[]>([])
</script>



