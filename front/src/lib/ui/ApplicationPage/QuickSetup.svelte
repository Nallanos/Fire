<script lang="ts">
  import { onMount } from "svelte";
  import { type App, getApp } from "$lib/api";
  import { page } from "$app/stores";
  import Input from "../Input.svelte";

  let app: App;
  onMount(async () => {
    app = await getApp($page.params?.id);
  });
</script>

<main>
  <h3 class="mx-auto container my-6 font-bold text-lg">Quick setup</h3>

  <div
    class="content-center mx-auto w-full container border-gray-800 border divide-y divide-gray-800"
  >
    <div
      class="flex flex-col full rounded-md gap-4 h-28 justify-center items-center"
    >
      <p class="font-bold">Write your code in our code editor !</p>
      <a href="{app?.id}/code-editor">
        <button
          class="border-gray-800 border h-[40px] w-[200px] rounded-md bg-white text-black"
        >
          Create your project now !
        </button>
      </a>
    </div>

    <div
      class="flex flex-col gap-2 w-full rounded-md h-28 justify-center items-center"
    >
      <p class="font-bold">... or import your code from your computer</p>

      <button class="border-gray-800 border h-[40px] w-[200px] rounded-md">
        Import code
      </button>
    </div>

    <div
      class="flex flex-col w-full rounded-md justify-center items-center gap-2 h-28"
    >
      <h3 class="font-bold">... or clone your code from a repository</h3>
      <Input />
    </div>

    <div
      class="flex flex-col w-full rounded-md justify-center items-center gap-2 h-28"
    >
      <h3 class="font-bold">... or clone your code from a docker directory</h3>
      <Input />
    </div>
  </div>
</main>
