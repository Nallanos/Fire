<script lang="ts">
  import { uploadRepo } from "$lib/api/uploadService/uploadRepo";
  import { inputValue } from "$lib/store";
  import { page } from "$app/stores";

  let id: string = $page.params?.id;
  let val = "";
  const updateVal = (event: Event) => {
    let inputElement = event.target as HTMLInputElement;
    val = inputElement.value;
    inputValue.set(val);
  };
  async function UploadRepo() {
    await uploadRepo({ repoURL: val, id: id });
  }
</script>

<form class="flex gap-2 justify-center" on:submit={UploadRepo}>
  <div
    class="flex py-4 items-center border border-gray-800 rounded-md sm:w-[80%] lg:w-[90%] w-[69%] h-[40px] p-4 mb-3"
  >
    <input
      on:input={updateVal}
      bind:value={val}
      type="text"
      placeholder="Example: github.com/MyUserName/MyRepository"
      class="placeholder-gray-500 text-gray-200 px-4 py-2 focus:outline-none bg-red bg-gray-950 w-full h-[30px]"
    />
  </div>
</form>
