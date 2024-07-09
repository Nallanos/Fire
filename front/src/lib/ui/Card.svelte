<script lang="ts">
  import type { App } from "$lib/api/apps/appTypes";
  export let app: App;
  let status = app.status;
  import StatusIcon from "./StatusIcon.svelte";
</script>

<a
  href={"/apps/" + app.id}
  class="bg-gray-950 rounded-lg flex flex-col border border-gray-800 p-6"
>
  <div class="flex items-center gap-2">
    <div class="text-xl font-bold">{app.name}</div>
    <div>
      <StatusIcon {status} />
    </div>
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
