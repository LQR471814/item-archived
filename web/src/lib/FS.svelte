<script lang="ts">
  import type { Archive } from "../archive";
  import { notifyError } from "./error";
  import { onMount } from "svelte";

  let {
    archive,
  }: {
    archive: Archive;
  } = $props();

  let cursor = $state<string[]>([]);

  type EntryList = {
    isItem: boolean;
    name: string;
  }[];

  let fs = $state<EntryList[]>([]);

  function select(path: string[]) {
    if (path.length > fs.length) {
    }
  }

  onMount(() => {
    archive.read([]).then((res) => {
      if (!res.children) {
        notifyError(new Error("Root entrylist contains null children."));
        return;
      }

      const entryList: EntryList = [];
      for (const itemName of res.children.itemNames) {
        entryList.push({ isItem: true, name: itemName });
      }
      for (const containerName of res.children.containerNames) {
        entryList.push({ isItem: false, name: containerName });
      }
      fs[0] = entryList;
    });
  });
</script>
