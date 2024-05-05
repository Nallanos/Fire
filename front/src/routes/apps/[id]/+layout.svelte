<script lang="ts">
  import { Settings } from "lucide-svelte";
  import { PanelsTopLeft } from "lucide-svelte";
  import { List } from "lucide-svelte";
  import { Braces } from "lucide-svelte";
  import { Computer } from "lucide-svelte";
  import HeaderIdApp from "$lib/ui/ApplicationPage/HeaderIdApp.svelte";
  import { page } from "$app/stores";
  let id: string = $page.params?.id;
</script>

<HeaderIdApp></HeaderIdApp>

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
