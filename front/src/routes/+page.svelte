<script lang="ts">
  import { onMount } from "svelte";
  let texts = ["Deploy Easily.", "Enjoy.", "Code."];
  let index = 0;
  let currentText = "Code.";

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

<main class="flex justify-center">
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
    <p class=" justify-self-start text-lg w-[640px] mr-auto">
      Stay in the flow with instant dev experiences. No more hours
      stashing/pulling/installing locally — just click, and start coding.
    </p>
  </div>
</main>
