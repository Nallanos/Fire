<script lang="ts">
  import { listApps } from "$lib/api/apps/listApps";
  import { inputValue } from "$lib/store";
  import { apps } from "$lib/store";
  import { Input } from "$lib/components/ui/input/index";
  let searchVal = "";
  const updateSearchVal = async (event: Event) => {
    let inputElement = event.target as HTMLInputElement;
    searchVal = inputElement.value;
    inputValue.set(searchVal);
    apps.set(await listApps());
  };
</script>

<div
  class="flex py-4 items-center rounded-md sm:w-[80%] lg:w-[90%] w-[70%] h-[40px] mb-3 transition-all duration-300 gap-2"
>
  <Input
    type="email"
    placeholder="Search for applications "
    class=" border border-gray-800 focus:outline-none"
    on:input={updateSearchVal}
    bind:value={searchVal}
  />
</div>
