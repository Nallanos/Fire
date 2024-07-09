<script lang="ts">
  import { apps } from "$lib/store";
  import { listApps } from "$lib/api/apps/listApps";
  import { Settings } from "lucide-svelte";
  import { PanelsTopLeft } from "lucide-svelte";
  import { List } from "lucide-svelte";
  import { Braces } from "lucide-svelte";
  import { Computer } from "lucide-svelte";
  import { page } from "$app/stores";
  import { deleteApp } from "$lib/api/apps/deleteApp";
  import type { LayoutData } from "./$types";
  export let data: LayoutData;

  let id: string = $page.params?.id;
  let app = data.app.data;
  console.log(data.app);
</script>

<header class="flex">
  <div
    class="bg-gray-950 rounded-lg py-6 flex justify-between items-center w-full container mx-auto"
  >
    <h1 class="text-2xl font-bold">{app?.name}</h1>

    <ul class="flex gap-4">
      <form action="?/deploy" method="POST">
        {#if app.status == "active"}
          <button
            class="border rounded-md border-gray-800 w-[80px] h-[40px] py-10px bg-white text-black"
            formaction="?/stopContainer"
          >
            Stop
          </button>
        {:else}
          <button
            class={`bg-white border rounded-md border-gray-800 w-[80px] h-[40px] py-10px text-black`}
          >
            Deploy
          </button>
        {/if}
      </form>
      <button
        class="bg-red-600 rounded-md w-[118px] h-[40px] py-10px"
        on:click={async () => {
          await deleteApp(id);
          apps.set(await listApps());
          window.location.href = "/apps";
        }}
      >
        Delete
      </button>
    </ul>
  </div>
</header>

<div class="flex h-24 py-6 container mx-auto w-full">
  <nav
    class="flex divide-x w-min border rounded border-gray-800 border-rounded h-12 items-center"
  >
    <a
      class="flex gap-2 items-center h-12 border-gray-800 px-14"
      href={`/apps/${id}`}
    >
      <PanelsTopLeft class="size-4" />
      Overview
    </a>

    <a
      class="flex h-12 px-6 border-gray-800 gap-2 items-center px-14"
      href={`/apps/${id}/env`}
    >
      <Braces class="size-4" />
      Env
    </a>

    <a
      class="h-12 px-6 border-gray-800 flex gap-2 items-center px-14"
      href={`/apps/${id}/logs`}
    >
      <List class="size-4" />
      Logs
    </a>

    <a
      class="h-12 px-6 border-gray-800 flex gap-2 items-center px-4 px-14"
      href={`/apps/${id}/domain`}
    >
      <Computer class="size-4" /> Domain
    </a>

    <a
      class="h-12 px-6 border-gray-800 flex items-center gap-2 px-14"
      href={`/apps/${id}/settings`}
    >
      <Settings class="size-4" />
      Settings
    </a>
  </nav>
</div>
<slot></slot>
