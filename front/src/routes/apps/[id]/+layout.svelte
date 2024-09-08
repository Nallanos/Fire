<script lang="ts">
  import {
    Settings,
    PanelsTopLeft,
    List,
    Braces,
    Computer,
    ChevronDown,
  } from "lucide-svelte";
  import { page } from "$app/stores";
  import { Button } from "$lib/components/ui/button/index";
  import type { LayoutData } from "./$types";
  import { onMount } from "svelte";

  export let data: LayoutData;
  let id: string = $page.params?.id;
  let app = data.app;
  let isOpen = false;
  let dropdownRef: HTMLElement | null = null;

  const links = [
    { href: `/apps/${id}`, icon: PanelsTopLeft, text: "Overview" },
    { href: `/apps/${id}/env`, icon: Braces, text: "Env" },
    { href: `/apps/${id}/logs`, icon: List, text: "Logs" },
    { href: `/apps/${id}/domain`, icon: Computer, text: "Domain" },
    { href: `/apps/${id}/settings`, icon: Settings, text: "Settings" },
  ];

  const toggleDropdown = () => {
    isOpen = !isOpen;
  };

  const handleClickOutside = (event: MouseEvent) => {
    if (dropdownRef && !dropdownRef.contains(event.target as Node)) {
      isOpen = false;
    }
  };

  onMount(() => {
    document.addEventListener("click", handleClickOutside);
    return () => {
      document.removeEventListener("click", handleClickOutside);
    };
  });
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
        <button
          formaction={app.status === "active" ? "?/stopContainer" : undefined}
          class="border rounded-md border-gray-800 w-[80px] h-[40px] py-10px bg-white text-black hover:bg-primary/90"
        >
          {app.status === "active" ? "Stop" : "Deploy"}
        </button>
      </form>
      <form action="?/deleteApp" method="post">
        <Button variant="destructive" class="bg-red-600" type="submit"
          >Delete</Button
        >
      </form>
    </ul>
  </div>
</header>

<div class="flex h-24 py-6 container mx-auto w-full relative md:hidden">
  <button
    bind:this={dropdownRef}
    class="flex items-center bg-[#09090b] hover:bg-[#27272a] border border-gray-800 gap-2 px-6 py-3 text-white rounded-md focus:outline-none"
    on:click={toggleDropdown}
  >
    Menu <ChevronDown class="size-4" />
  </button>

  {#if isOpen}
    <nav
      class="absolute bg-[#09090b] border border-gray-800 text-white mt-2 rounded shadow-lg z-10 w-56"
    >
      {#each links as { href, icon: Icon, text }}
        <a
          {href}
          class="flex gap-2 items-center h-12 px-4 hover:bg-[#27272a] w-full"
        >
          <Icon class="size-4" />
          {text}
        </a>
      {/each}
    </nav>
  {/if}
</div>

<div class="hidden md:flex h-24 py-6 container mx-auto w-full">
  <nav class="flex divide-x border border-gray-800 rounded items-center">
    {#each links as { href, icon: Icon, text }}
      <a
        {href}
        class="flex h-12 px-6 border-gray-800 gap-2 items-center px-14 hover:bg-[#27272a]"
      >
        <Icon class="size-4" />
        {text}
      </a>
    {/each}
  </nav>
</div>

<slot />
