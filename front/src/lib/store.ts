import { writable } from "svelte/store";
import type { SearchedImage } from "lib";
export const imagesDockerHub = writable<SearchedImage[]>([]);

