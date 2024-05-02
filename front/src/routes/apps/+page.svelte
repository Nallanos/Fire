<script lang="ts">
  import { onMount } from "svelte";
  import { writable } from "svelte/store";
  import Input from "$lib/ui/Apps-input.svelte";
  import Card from "$lib/ui/Card.svelte";
  import { listApps, type App } from "$lib/api";
  import { inputValue } from "$lib/store";

  let apps = writable<App[]>([]);
  let displayedApps: App[] = [];

  onMount(async () => {
    const initialApps: App[] = await listApps();
    apps.set(initialApps);

    apps.subscribe((val: App[]) => {
      displayedApps = val;
      inputValue.subscribe((inputVal: string) => {
        let modifiedVal = inputVal.toLowerCase().replace(/\s/g, "");
        if (inputVal === "" || inputVal === null || inputVal === undefined) {
          displayedApps = val;
        } else {
          displayedApps = val.filter((item: App) =>
            item.name.toLowerCase().includes(modifiedVal)
          );
        }
      });
    });
  });
</script>

<div class="container mx-auto px-4">
  <h1 class="text-3xl font-bold py-8 pr-6">Apps</h1>
  <Input />
  <div class="grid lg:grid-cols-3 gap-6 grid-cols-1">
    {#if displayedApps != null}
      {#each displayedApps as app}
        <Card {app} />
      {/each}
    {/if}
  </div>
</div>
