import { notifyError } from "./error";

export const ITEM = "item";
export const CONTAINER = "container";

export function parseName(name: string): {
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

export type EntryList = {
  isItem: boolean;
  name: string;
}[];

export function entryList(children?: {
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
