<script>
  import "../app.css";
  import { ModeWatcher } from "mode-watcher";
  import Sun from "lucide-svelte/icons/sun";
  import Moon from "lucide-svelte/icons/moon";
  import { toggleMode } from "mode-watcher";
  import { Button } from "$lib/components/ui/button/index.ts";
  const items = [
    { name: "Home", href: "/" },
    { name: "Services", href: "/services" },
    { name: "Apps", href: "/apps" },
  ];
</script>

<main class="text-white bg-gray-950 min-h-screen">
  <header class="border-b border-gray-800 h-20">
    <div
      class="flex items-center container mx-auto lg:mx-auto lg:justify-evenly md: mx-auto md:justify-evenly min-[450px]:justify-evenly min-[450px]:mx-auto pr-4 gap-10 h-full"
    >
      <a href="/" class="w-[96px] h-[96px] flex">
        <img src="/Fire.png" alt="Fire logo" class="w-[96px] h-[96px]" />
      </a>

      <ul class="flex lg:gap-20 gap-8 text-md">
        {#each items as item}
          <li>
            <a class="hover:text-gray-300" href={item.href}>
              {item.name}
            </a>
          </li>
        {/each}
      </ul>

      <Button on:click={toggleMode} variant="secondary" size="icon">
        <Sun
          class="h-[1.2rem] w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
        />
        <Moon
          class="absolute h-[1.2rem] w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
        />
        <span class="sr-only">Toggle theme</span>
      </Button>
    </div>
  </header>
  <ModeWatcher />
  <div>
    <slot />
  </div>
</main>
