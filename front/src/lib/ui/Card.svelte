<script lang="ts">
  import type { App } from "$lib/api";
  export let app: App;
</script>

<a
  href={"/apps/" + app.id}
  class="bg-gray-950 rounded-lg flex flex-col border border-gray-800 p-6"
>
  <div class="flex items-center">
    <div class="text-xl font-bold">{app.name}</div>
    <div class="bg-red-500 rounded-full w-[10px] h-[10px] ml-auto"></div>
  </div>
  <ul class="flex flex-col">
    <li class="flex gap-3">
      {app.id}
    </li>
    <li class="text-gray-200">
      Listening on port
      <span>{app.port}</span>
    </li>
  </ul>
</a>
