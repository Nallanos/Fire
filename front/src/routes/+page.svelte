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

    <div class="flex gap-6">
      <Button href="/sign-up">Sign up</Button>
      <Button href="/login" variant="outline">Log in</Button>
    </div>
  </div>
</header>
<div
  class="container items-center justify-center flex flex-col w-2/3 justify-self-center"
>
  <div class="flex flex-col container py-20">
    <h1 class="text-7xl font-bold font-mono text-blue-500">Click.</h1>
    <span class="h-[72px]">
      <h1 class="text-7xl font-bold font-mono blur-out blur-in">
        {currentText}
      </h1>
    </span>
    <h1 class="text-red-500 text-7xl font-bold font-mono">Done.</h1>
  </div>
  <p class=" justify-self-start text-lg lg:w-[640px] mr-auto">
    Stay in the flow with instant dev experiences. No more hours
    stashing/pulling/installing locally — just click, and start coding.
  </p>
</div>
