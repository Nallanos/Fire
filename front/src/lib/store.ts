import { writable } from "svelte/store";
import { type App } from "./api";
export const inputValue = writable("");
export let apps = writable<App[]>([]);
