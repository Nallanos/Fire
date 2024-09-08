<script lang="ts">
  import { onMount } from "svelte";
  import Sun from "lucide-svelte/icons/sun";
  import Moon from "lucide-svelte/icons/moon";
  import { toggleMode } from "mode-watcher";
  import { Button } from "$lib/components/ui/button/index.ts";
  let texts = ["Deploy.", "Enjoy.", "Code."];
  let index = 0;
  let currentText = "Code.";
  const items = [
    { name: "Home", href: "/" },
    { name: "Pricing", href: "/princings" },
    { name: "Apps", href: "/apps" },
  ];
  onMount(() => {
    const interval = setInterval(() => {
      eraseWriter(() => {
        index = (index + 1) % texts.length;
        typeWriter(texts[index]);
      });
    }, 3500);

    return () => {
      clearInterval(interval);
    };
  });

  function typeWriter(text: string, i = 0) {
    if (i < text.length) {
      currentText += text.charAt(i);
      setTimeout(() => typeWriter(text, i + 1), 100);
    }
  }

  function eraseWriter(callback: () => void, i = currentText.length) {
    if (i > 0) {
      currentText = currentText.slice(0, i - 1);
      setTimeout(() => eraseWriter(callback, i - 1), 160);
    } else {
      callback();
    }
  }
</script>

<header class="border-b border-gray-800 h-20 md:justify-evenly w-full">
  <div class="flex items-center container w-full md:justify-evenly h-full">
    <a href="/" class="flex">
      <img src="/Fire.png" alt="Fire logo" class="w-[96px] h-[96px]" />
    </a>

    <div class="text-md gap-10 hidden md:flex">
      {#each items as item}
        <a class="hover:text-gray-300" href={item.href}>
          {item.name}
        </a>
      {/each}
    </div>

    <div class="flex gap-6 md:gap-4 justify-evenly md:justify-end w-2/4">
      <Button href="/sign-up" class="w-[60px] ">Sign up</Button>
      <Button href="/login" variant="outline" class="w-[60px]">Log in</Button>
      <Button
        on:click={toggleMode}
        variant="secondary"
        size="icon"
        class="w-[30px]"
      >
        <Sun
          class=" md:w-[1.2rem] rotate-0 scale-100 transition-all dark:-rotate-90 dark:scale-0"
        />
        <Moon
          class="absolute h-[1.2rem] w-[0.9rem] md:w-[1.2rem] rotate-90 scale-0 transition-all dark:rotate-0 dark:scale-100"
        />
        <span class="sr-only">Toggle theme</span>
      </Button>
    </div>
  </div>
</header>

<main class="container items-center justify-center flex flex-col">
  <div class="flex flex-col py-20">
    <h1 class="md:text-7xl text-5xl font-bold font-mono text-blue-500">
      Click.
    </h1>
    <span>
      <h1
        class="md:text-7xl text-5xl font-bold font-mono blur-out blur-in h-[48px] md:h-[72px]"
      >
        {currentText}
      </h1>
    </span>
    <h1 class="text-red-500 text-5xl md:text-7xl font-bold font-mono">Done.</h1>

    <p class="pt-8">
      Stay in the flow with instant dev experiences. No more hours
      stashing/pulling/installing locally — just click, and start coding.
    </p>
  </div>
</main>
