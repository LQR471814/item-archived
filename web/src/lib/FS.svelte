<script lang="ts">
  import type { Archive } from "../archive";
  import type { EntryMetadata } from "../api/v1/api_pb";
  import { notifyError } from "./error";
  import { onMount, tick } from "svelte";
  import { twMerge } from "tailwind-merge";
  import Image from "./Image.svelte";
  import { type EntryList, entryList, parseName, CONTAINER, ITEM } from "./fs"

  let {
    archive,
  }: {
    archive: Archive;
  } = $props();

  let cursor = $state<string[]>([]);
  let fs = $state<EntryList[]>([]);
  let meta = $state<EntryMetadata>();
  let selectedEntry = $state<HTMLDivElement>();
  let draggedOver = $state<number>();

  $effect(() => {
    // this is to tell the svelte compiler that this $effect depends on selectedEntry
    selectedEntry;
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
      "flex gap-1 w-full items-center",
      selected
        ? "bg-blue-500 hover:bg-blue-500 text-zinc-50 font-semibold"
        : "",
    )}
    {onclick}
    draggable="true"
  >
    {#if type === CONTAINER}
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="size-5"
        viewBox="0 0 24 24"
        fill="currentColor"
        ><path
          d="M22 20V7L20 3H4L2 7.00353V20C2 20.5523 2.44772 21 3 21H21C21.5523 21 22 20.5523 22 20ZM4 9H20V19H4V9ZM5.236 5H18.764L19.764 7H4.237L5.236 5ZM15 11H9V13H15V11Z"
        ></path>
      </svg>
    {:else if type === ITEM}
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="size-5"
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

<div class="flex flex-col border-t border-b border-solid border-zinc-300">
  <div class="flex gap-3 overflow-x-auto">
    {#each fs as entryList, i}
      {#if entryList.length > 0}
        <div
          class={twMerge(
            "flex flex-col border-l border-r border-solid border-zinc-300",
            draggedOver === i ? "bg-zinc-300" : "",
          )}
          role="region"
          ondragover={(e) => {
            e.preventDefault();
            if (draggedOver === i) {
              return;
            }
            draggedOver = i;
          }}
          ondragleave={(e) => {
            e.preventDefault();
            draggedOver = undefined;
          }}
          ondragend={(e) => {
            e.preventDefault();
            draggedOver = undefined;
          }}
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
      class="flex flex-col gap-x-3 border-t border-l border-r border-solid border-zinc-300 flex-wrap"
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
        {#if meta.tags.length > 0}
          {#each meta.tags as tag}
            <div class="border border-solid border-zinc-300 w-fit h-fit">
              {tag}
            </div>
          {/each}
        {:else}
          <p>-</p>
        {/if}
      </div>

      {#if meta.image !== undefined && meta.imageFormat !== undefined}
        <div class="flex gap-2">
          <p class="text-zinc-500 font-mono">Image:</p>
          <Image
            className="max-h-40"
            bytes={meta.image.buffer as ArrayBuffer}
            format={meta.imageFormat}
            alt="image"
          />
        </div>
      {/if}

      <button class="flex gap-1 items-center w-fit">
        <svg
          class="size-5"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="currentColor"
          ><path
            d="M15.7279 9.57627L14.3137 8.16206L5 17.4758V18.89H6.41421L15.7279 9.57627ZM17.1421 8.16206L18.5563 6.74785L17.1421 5.33363L15.7279 6.74785L17.1421 8.16206ZM7.24264 20.89H3V16.6473L16.435 3.21231C16.8256 2.82179 17.4587 2.82179 17.8492 3.21231L20.6777 6.04074C21.0682 6.43126 21.0682 7.06443 20.6777 7.45495L7.24264 20.89Z"
          ></path>
        </svg>
        <span>Edit</span>
      </button>

      <button class="flex gap-1 items-center text-red-500 w-fit">
        <svg
          class="size-5"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          fill="currentColor"
        >
          <path
            d="M17 6H22V8H20V21C20 21.5523 19.5523 22 19 22H5C4.44772 22 4 21.5523 4 21V8H2V6H7V3C7 2.44772 7.44772 2 8 2H16C16.5523 2 17 2.44772 17 3V6ZM18 8H6V20H18V8ZM9 11H11V17H9V11ZM13 11H15V17H13V11ZM9 4V6H15V4H9Z"
          ></path>
        </svg>
        <span>Delete</span>
      </button>
    </div>
  {/if}
</div>
