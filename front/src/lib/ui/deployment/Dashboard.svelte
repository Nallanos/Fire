<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    import type { MetricData } from "lib/index";
    import Courbe from "./Courbe.svelte";
    import type { App } from "lib/api/apps/appTypes";

    export let app: App;
    const app_id = app.id;
    export let token: string;

    let socket: WebSocket;
    let stats: MetricData | null = null;
    let chartComponent: Courbe;

    function connectWebSocket() {
        socket = new WebSocket(
            `ws://74.235.166.230:8081/${app_id}/stats?token=${encodeURIComponent(token)}`,
        );

        socket.addEventListener("open", () => {
            console.log("WebSocket Dashboard connection established");
        });

        socket.addEventListener("close", () => {
            console.log("WebSocket Dashboard connection closed");
        });

        socket.onmessage = (event) => {
            stats = JSON.parse(event.data);
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

{#if stats && app.status == "active"}
    <Courbe bind:this={chartComponent} MetricData={stats} />
{/if}
