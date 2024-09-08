<script>
    import "../app.css";
    import { ModeWatcher } from "mode-watcher";
</script>

<main class="text-white bg-black min-h-screen">
    <div>
        <slot />
    </div>
    <ModeWatcher />
</main>
