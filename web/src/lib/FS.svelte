<script lang="ts">
  import type { Archive } from "../archive";
  import type { EntryMetadata } from "../api/v1/api_pb";
  import { notifyError } from "./error";
  import { onMount, tick } from "svelte";
  import { twMerge } from "tailwind-merge";

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
  let meta = $state<EntryMetadata>();

  function entryList(children?: {
    itemNames: string[];
    containerNames: string[];
  }): EntryList {
    if (!children) {
      return [];
    }
    const entryList: EntryList = [];
    for (const itemName of children.itemNames) {
      entryList.push({ isItem: true, name: itemName });
    }
    for (const containerName of children.containerNames) {
      entryList.push({ isItem: false, name: containerName });
    }
    return entryList;
  }

  const ITEM = "item";
  const CONTAINER = "container";

  function parseName(name: string): {
    id: string;
    tags: string[];
    type: typeof ITEM | typeof CONTAINER;
  } {
    const segments = name.split(".");
    if (segments.length === 0) {
      notifyError(new Error(`invalid name: '${name}'`));
      return { id: name, tags: [], type: ITEM };
    }
    return {
      id: segments[0],
      tags: segments.slice(1, segments.length - 1),
      type: segments[segments.length - 1] === ITEM ? ITEM : CONTAINER,
    };
  }

  let selectedEntry = $state<HTMLDivElement>();

  $effect(() => {
    // this is to tell the svelte compiler that this $effect depends on selectedEntry
    selectedEntry
    tick().then(() => {
      if (!selectedEntry) {
        return;
      }
      selectedEntry.scrollIntoView();
    });
  });

  async function select(path: string[]) {
    cursor = path;

    if (path.length > fs.length) {
      notifyError(
        new Error(
          `selected path ${path} that is outside the current fs range.`,
        ),
      );
      return;
    }

    const { metadata, children } = await archive.read(path);
    meta = metadata;
    fs = [...fs.slice(0, path.length), entryList(children)];
  }

  onMount(() => {
    archive
      .read([])
      .then((res) => {
        if (!res.children) {
          notifyError(new Error("Root entrylist contains null children."));
          return;
        }
        fs[0] = entryList(res.children);
      })
      .catch((err) => {
        notifyError(err);
      });
  });
</script>

{#snippet entryButton(
  id: string,
  type: typeof CONTAINER | typeof ITEM,
  selected: boolean,
  onclick: () => void,
)}
  <button
    class={twMerge(
      "text-md hover:cursor-default hover:bg-zinc-300 px-3 text-left",
      "flex gap-1 w-full",
      selected
        ? "bg-blue-500 hover:bg-blue-500 text-zinc-50 font-semibold"
        : "",
    )}
    {onclick}
  >
    {#if type === CONTAINER}
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="min-w-6 min-h-6 max-w-6 max-h-6"
        viewBox="0 0 24 24"
        fill="currentColor"
      >
        <path
          d="M3 10H2V4.00293C2 3.44903 2.45531 3 2.9918 3H21.0082C21.556 3 22 3.43788 22 4.00293V10H21V20.0015C21 20.553 20.5551 21 20.0066 21H3.9934C3.44476 21 3 20.5525 3 20.0015V10ZM19 10H5V19H19V10ZM4 5V8H20V5H4ZM9 12H15V14H9V12Z"
        ></path>
      </svg>
    {:else if type === ITEM}
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="min-w-6 min-h-6 max-w-6 max-h-6"
        viewBox="0 0 24 24"
        fill="currentColor"
      >
        <path
          d="M12 1L21.5 6.5V17.5L12 23L2.5 17.5V6.5L12 1ZM5.49388 7.0777L13.0001 11.4234V20.11L19.5 16.3469V7.65311L12 3.311L5.49388 7.0777ZM4.5 8.81329V16.3469L11.0001 20.1101V12.5765L4.5 8.81329Z"
        ></path>
      </svg>
    {/if}

    {id}
  </button>
{/snippet}

<div class="flex flex-col">
  <div class="flex gap-3 overflow-x-auto">
    {#each fs as entryList, i}
      {#if entryList.length > 0}
        <div
          class="flex flex-col border-l border-r border-solid border-zinc-300"
        >
          {#each entryList as entry}
            {@const parsedName = parseName(entry.name)}
            {@const selected = cursor[i] === entry.name}
            {@const onclick = () => {
              select([...cursor.slice(0, i), entry.name]);
            }}

            {#if selected}
              <div class="w-full" bind:this={selectedEntry}>
                {@render entryButton(
                  parsedName.id,
                  parsedName.type,
                  selected,
                  onclick,
                )}
              </div>
            {:else}
              {@render entryButton(
                parsedName.id,
                parsedName.type,
                selected,
                onclick,
              )}
            {/if}
          {/each}
        </div>
      {/if}
    {/each}
  </div>

  {#if meta}
    <div
      class="flex gap-x-3 gap-y-0 justify-evenly border border-solid border-zinc-300 flex-wrap"
    >
      <div class="flex gap-2">
        <p class="text-zinc-500 font-mono">ID:</p>
        <p>{meta.id}</p>
      </div>

      <div class="flex gap-2">
        <p class="text-zinc-500 font-mono">Description:</p>
        <p>{meta.description ?? "-"}</p>
      </div>

      <div class="flex gap-2">
        <p class="text-zinc-500 font-mono">Tags:</p>
        {#each meta.tags as tag}
          <div class="border border-solid border-zinc-300">
            {tag}
          </div>
        {/each}
      </div>
    </div>
  {/if}
</div>
