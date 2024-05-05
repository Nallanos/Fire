<script lang="ts">
  export let status: string;
  let backgroundColor = "bg-[#262f3b]";
  let iconColor = "bg-[#9ca3af]";
  $: {
    switch (status) {
      case "completed":
        backgroundColor = "bg-[#192c2f]";
        iconColor = "bg-[#00bfa5]";
        break;
      case "inactive":
        backgroundColor = "bg-[#262f3b]";
        iconColor = "bg-[#9ca3af]";
        break;
      case "failed":
        backgroundColor = "bg-[#2a2230]";
        iconColor = "bg-[#ef4444]";
        break;
    }
  }
</script>

<div class="">
  <div class="flex-none relative">
    <div class={`rounded-full p-1 ${backgroundColor}`}>
      <div class={`h-2 w-2 rounded-full ${iconColor}`}></div>
    </div>
  </div>
</div>
