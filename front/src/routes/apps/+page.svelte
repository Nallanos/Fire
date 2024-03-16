<script lang="ts">
  import { onMount } from "svelte";
  import Input from "$lib/ui/Apps-input.svelte";
  import Card from "$lib/ui/Card.svelte";
  import { listApps, type App } from "$lib/api";
  import { inputValue } from "$lib/store";
  let apps: App[] = [];
  let displayedApps: App[] = apps;

  onMount(async () => {
    apps = await listApps();
    displayedApps = apps;

    inputValue.subscribe((val: string) => {
      if (val === "" || val === null || val === undefined) {
        displayedApps = apps;
      } else {
        displayedApps = apps.filter((item: App) => item.name.includes(val));
      }
    });
  });
</script>

<div>
  <h1 class="text-3xl font-bold py-8 pr-6">Apps</h1>
  <Input />
  <div class="grid lg:grid-cols-3 gap-6 grid-cols-1">
    {#each displayedApps as app}
      <Card {app} />
    {/each}
  </div>
</div>
