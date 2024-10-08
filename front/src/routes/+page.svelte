<script lang="ts">
  import { onMount } from "svelte";
  import HomeCard from "lib/ui/HomeCard.svelte";
  let texts = ["Deploy.", "Enjoy.", "Code."];
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

<main class="container items-center justify-center flex w-full flex-col">
  <div class="flex flex-col pt-20 text-center">
    <h1
      class="md:text-7xl text-5xl font-bold text-blue-500 font-sans transition-all duration-500 ease-in-out"
    >
      Click.
    </h1>

    <h1
      class="md:text-7xl text-5xl font-bold font-sans text-white md:h-[72px] h-[48px] transition-colors duration-500 ease-in-out"
    >
      {currentText}
    </h1>

    <h1
      class="text-red-500 text-5xl md:text-7xl font-bold font-sans transition-all duration-500 ease-in-out"
    >
      Done.
    </h1>

    <p
      class="pt-8 max-w-2xl mx-auto text-lg text-gray-400 leading-relaxed font-light"
    >
      Stay in the flow with instant dev experiences. No more hours
      stashing/pulling/installing locally — just click, and start coding.
    </p>
  </div>

  <!-- Carte animée -->
  <div class="w-full px-4 sm:w-3/4 lg:w-1/2 mt-12">
    <HomeCard />
  </div>
</main>

<style global>
  /* Ajoute des animations subtiles */
  @keyframes fadeIn {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }

  h1 {
    animation: fadeIn 1s ease-in-out;
  }

  @keyframes blinkCursor {
    from {
      border-right-color: rgba(255, 255, 255, 0.75);
    }
    to {
      border-right-color: transparent;
    }
  }
</style>
