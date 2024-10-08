<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    import type { MetricData } from "lib/index";
    import type { App } from "lib/api/apps/appTypes";

    export let app: App;
    export let token: string;

    const app_id = app.id;
    let socket: WebSocket;
    let logs: string | null = null;

    function connectWebSocket() {
        socket = new WebSocket(
            `ws://74.235.166.230:8081/${app_id}/logs?token=${encodeURIComponent(token)}`,
        );

        socket.addEventListener("open", () => {
            console.log("WebSocket logs connection established");
        });

        socket.addEventListener("close", () => {
            console.log("WebSocket logs connection closed");
        });

        socket.onmessage = (event) => {
            logs = event.data;
        };
    }

    function disconnectWebSocket() {
        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.close(1000, "web socket connections closed on destroy");
        }
    }

    onMount(() => {
        if (app.status == "active") {
            connectWebSocket();
        }
    });

    onDestroy(() => {
        disconnectWebSocket();
    });
</script>

{#if logs}
    <div
        class="mt-6 bg-neutral-900 mx-auto text-white font-mono rounded-md shadow-md overflow-auto h-[620px] container"
    >
        {@html logs.replace(/\n/g, "<br>")}
    </div>
{/if}
