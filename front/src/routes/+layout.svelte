<script>
    import "../app.css";
    import { ModeWatcher } from "mode-watcher";
    import { Button } from "$lib/components/ui/button/index";
    import Sun from "lucide-svelte/icons/sun";
    import Moon from "lucide-svelte/icons/moon";
    import { toggleMode } from "mode-watcher";

    const items = [
        { name: "Home", href: "/" },
        { name: "Pricing", href: "/pricing" },
        { name: "Apps", href: "/apps" },
    ];
</script>

<main
    class="text-white bg-gradient-to-b from-black from-60% to-neutral-900 min-h-screen"
>
    <header
        class="flex justify-between items-center h-16 -px-6 border-b border-gray-800"
    >
        <div class="flex items-center space-x-4">
            <a href="/" class="flex">
                <img
                    src="/Fire.png"
                    alt="Fire logo"
                    class="w-[64px] h-[64px]"
                />
            </a>
            <nav class="hidden md:flex space-x-8">
                {#each items as item}
                    <a
                        class="hover:text-gray-300 font-semibold tracking-wide"
                        href={item.href}
                    >
                        {item.name}
                    </a>
                {/each}
            </nav>
            
        </div>

        <div class="flex items-center space-x-4">
            <Button href="/sign-up" class="text-sm font-semibold"
                >Sign up</Button
            >
            <Button
                href="/login"
                variant="outline"
                class="text-sm font-semibold">Log in</Button
            >
            <Button
                on:click={toggleMode}
                variant="secondary"
                size="icon"
                class="relative w-10 h-10 flex items-center justify-center"
            >
                <Sun
                    class="w-5 h-5 transition-transform transform dark:rotate-0 dark:scale-100 scale-0"
                />
                <Moon
                    class="absolute w-5 h-5 transition-transform transform dark:rotate-90 dark:scale-100 scale-0"
                />
                <span class="sr-only">Toggle theme</span>
            </Button>
        </div>
    </header>
    <div>
        <slot />
    </div>
    <ModeWatcher />
</main>
