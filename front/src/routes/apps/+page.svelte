<script lang="ts">
  import type { PageServerData } from "./$types";
  import { Input } from "$lib/components/ui/input";
  import Card from "$lib/ui/Card.svelte";
  import type { App } from "$lib/api/apps/appTypes";
  import { ExternalLink } from "lucide-svelte";
  import { Button } from "$lib/components/ui/button/index";

  let displayedApps: App[] = [];
  let inputVal = "";
  export let data: PageServerData;

  $: displayedApps = data.apps
    ? data.apps.filter((app: App) =>
        app.name
          .toLowerCase()
          .includes(inputVal.toLowerCase().replace(/\s/g, "")),
      )
    : [];
</script>

<div class="container mx-auto px-4">
  <h2 class="text-2xl font-bold pb-6 pr-6">Apps</h2>
  <div class="flex gap-2">
    <div
      class="flex py-4 items-center rounded-md sm:w-[80%] lg:w-[90%] w-[70%] h-[40px] mb-3 transition-all duration-300 gap-2"
    >
      <Input
        id="inputVal"
        name="inputVal"
        placeholder="Search for applications"
        class="border border-gray-800 focus:outline-none"
        bind:value={inputVal}
      />
    </div>

    <Button href="/apps/create" variant="secondary" class="gap-2">
      <ExternalLink class="size-4" />
      Add a new...
    </Button>
  </div>

  <div class="grid lg:grid-cols-3 gap-6 grid-cols-1">
    {#each displayedApps as app}
      <Card {app} />
    {/each}
  </div>
</div>
