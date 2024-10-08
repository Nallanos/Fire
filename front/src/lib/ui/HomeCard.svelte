<script>
    import { Sun, Cpu, GitMerge } from "lucide-svelte";
    const cardStyle =
        "p-6 bg-neutral-900 rounded-lg border border-gray-700 shadow-md";
</script>

<section
    class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 py-10 text-center"
>
    <div class={cardStyle}>
        <GitMerge class="w-10 h-10 mx-auto text-white mb-4" />
        <h2 class="text-lg font-bold text-white">Open-Source</h2>
        <p class="text-gray-400">
            Fire <strong>maintains and contributes</strong> to open-source projects.
        </p>
    </div>

    <div class={cardStyle}>
        <Cpu class="w-10 h-10 mx-auto text-white mb-4" />
        <h2 class="text-lg font-bold text-white">Subsecond Boot Time</h2>
        <p class="text-gray-400">Your machines boot in 300ms or less.</p>
    </div>

    <div class={cardStyle}>
        <Sun class="w-10 h-10 mx-auto text-white mb-4" />
        <h2 class="text-lg font-bold text-white">Sustainability</h2>
        <p class="text-gray-400">
            Optimized resources for eco-friendly deployments.
        </p>
    </div>
</section>
