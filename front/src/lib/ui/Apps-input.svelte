<script lang="ts">
  import type { App } from "./../api/apps";
  import { createApp } from "$lib/api/apps";
  import { inputValue, apps } from "$lib/store";
  let val = "";
  const updateVal = (event: Event) => {
    let inputElement = event.target as HTMLInputElement;
    val = inputElement.value;
    inputValue.set(val);
  };
  async function createAppWithSearchTerm() {
    const newApp: App = await createApp({ name: val });
    apps.update((currentApps) => [...currentApps, newApp]);
  }
</script>

<form class="flex gap-2 justify-center" on:submit={createAppWithSearchTerm}>
  <div
    class="flex py-4 items-center border border-gray-800 rounded-md sm:w-[80%] lg:w-[90%] w-[69%] h-[40px] p-4 mb-3"
  >
    <span data-geist-input-prefix=""
      ><svg
        data-testid="geist-icon"
        height="16"
        stroke-linejoin="round"
        viewBox="0 0 16 16"
        width="16"
        style="width: 16px; height: 16px; color: currentcolor;"
        ><path
          fill-rule="evenodd"
          clip-rule="evenodd"
          d="M1.5 6.5C1.5 3.73858 3.73858 1.5 6.5 1.5C9.26142 1.5 11.5 3.73858 11.5 6.5C11.5 9.26142 9.26142 11.5 6.5 11.5C3.73858 11.5 1.5 9.26142 1.5 6.5ZM6.5 0C2.91015 0 0 2.91015 0 6.5C0 10.0899 2.91015 13 6.5 13C8.02469 13 9.42677 12.475 10.5353 11.596L13.9697 15.0303L14.5 15.5607L15.5607 14.5L15.0303 13.9697L11.596 10.5353C12.475 9.42677 13 8.02469 13 6.5C13 2.91015 10.0899 0 6.5 0Z"
          fill="currentColor"
        ></path></svg
      >
    </span>
    <input
      required
      on:input={updateVal}
      bind:value={val}
      type="text"
      placeholder="Search Repositories or Create a new one..."
      class="placeholder-gray-500 text-gray-200 px-4 py-2 focus:outline-none bg-red bg-gray-950 w-full h-[30px]"
    />
  </div>
  <a
    href="/apps"
    class="bg-white rounded-md h-[40px] lg:w-[10%] sm:w-[39%] md:w-[20%] text-black flex items-center justify-center"
  >
    <button type="submit" on:click={createAppWithSearchTerm}>Add project</button
    ></a
  >
</form>
