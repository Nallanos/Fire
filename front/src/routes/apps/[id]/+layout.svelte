<script lang="ts">
  import { Settings } from "lucide-svelte";
  import { PanelsTopLeft } from "lucide-svelte";
  import { List } from "lucide-svelte";
  import { Braces } from "lucide-svelte";
  import { Computer } from "lucide-svelte";
  import { page } from "$app/stores";
  import type { LayoutData } from "./$types";
  import { Button } from "$lib/components/ui/button/index";

  export let data: LayoutData;

  let id: string = $page.params?.id;
  let app = data.app;
</script>

<header class="flex">
  <div
    class="bg-black rounded-lg py-6 flex justify-between items-center w-full container mx-auto"
  >
    <div class="flex flex-col gap-2">
      <h1 class="text-2xl font-bold">{app?.name}</h1>
      <p class="text-xs">{app?.image}</p>
    </div>

    <ul class="flex gap-4">
      <form action="?/deploy" method="POST">
        {#if app.status == "active"}
          <button
            formaction="?/stopContainer"
            class="border rounded-md border-gray-800 w-[80px] h-[40px] py-10px bg-white text-black"
          >
            Stop
          </button>
        {:else}
          <button
            class={`bg-white border rounded-md border-gray-800 w-[80px] h-[40px] py-10px text-black hover:bg-primary/90`}
          >
            Deploy
          </button>
        {/if}
      </form>
      <form action="?/deleteApp" method="post">
        <Button variant="destructive" class="bg-red-600 " type="submit"
          >Delete</Button
        >
      </form>
    </ul>
  </div>
</header>

<div class="flex h-24 py-6 container mx-auto w-full">
  <nav
    class="flex divide-x w-min border rounded border-gray-800 border-rounded h-12 items-center"
  >
    <a
      class="flex gap-2 items-center h-12 border-gray-800 px-14 hover:bg-[#27272a]"
      href={`/apps/${id}`}
    >
      <PanelsTopLeft class="size-4" />
      Overview
    </a>

    <a
      class="flex h-12 px-6 border-gray-800 gap-2 items-center px-14 hover:bg-[#27272a]"
      href={`/apps/${id}/env`}
    >
      <Braces class="size-4" />
      Env
    </a>

    <a
      class="h-12 px-6 border-gray-800 flex gap-2 items-center px-14 hover:bg-[#27272a]"
      href={`/apps/${id}/logs`}
    >
      <List class="size-4" />
      Logs
    </a>

    <a
      class="h-12 px-6 border-gray-800 flex gap-2 items-center px-4 px-14 hover:bg-[#27272a]"
      href={`/apps/${id}/domain`}
    >
      <Computer class="size-4" /> Domain
    </a>

    <a
      class="h-12 px-6 border-gray-800 flex items-center gap-2 px-14 hover:bg-[#27272a]"
      href={`/apps/${id}/settings`}
    >
      <Settings class="size-4" />
      Settings
    </a>
  </nav>
</div>

<slot />
