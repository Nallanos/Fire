<script>
    import "../app.css";
    import { ModeWatcher } from "mode-watcher";
    import Sun from "lucide-svelte/icons/sun";
    import Moon from "lucide-svelte/icons/moon";
    import { toggleMode } from "mode-watcher";
    import { Button } from "$lib/components/ui/button/index.ts";
    const items = [
        { name: "Home", href: "/" },
        { name: "Princing", href: "/pricing" },
    ];
</script>

<main class="text-white bg-black min-h-screen">
    <div>
        <slot />
    </div>
    <ModeWatcher />
</main>
