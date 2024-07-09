<script lang="ts">
  import type { PageServerData } from "./$types";
  import { onMount } from "svelte";
  import Input from "$lib/ui/Apps-input.svelte";
  import Card from "$lib/ui/Card.svelte";
  import type { App } from "$lib/api/apps/appTypes";
  import { inputValue, apps } from "$lib/store";
  let displayedApps: App[] = [];
  export let data: PageServerData;
  onMount(async () => {
    apps.set(data.apps);
    apps.subscribe((val: App[]) => {
      displayedApps = val;
      inputValue.subscribe((inputVal: string) => {
        let modifiedVal = inputVal.toLowerCase().replace(/\s/g, "");
        if (
          inputVal === "" ||
          inputVal === null ||
          inputVal === undefined ||
          val === null
        ) {
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
