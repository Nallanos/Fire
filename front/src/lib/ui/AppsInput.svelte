<script lang="ts">
  import { listApps } from "$lib/api/apps/listApps";
  import { inputValue } from "$lib/store";
  import { Input } from "$lib/components/ui/input/index";
  let searchVal = "";
  const updateSearchVal = async (event: Event) => {
    let inputElement = event.target as HTMLInputElement;
    searchVal = inputElement.value;
    inputValue.set(searchVal);
  };
</script>

