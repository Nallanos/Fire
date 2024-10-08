<script lang="ts">
    import { Button } from "$lib/components/ui/button/index.js";
    import Check from "lucide-svelte/icons/check";
    import ChevronsUpDown from "lucide-svelte/icons/chevrons-up-down";
    import { Input } from "$lib/components/ui/input/index.js";
    import * as Popover from "$lib/components/ui/popover/index.js";
    import * as Command from "$lib/components/ui/command/index.js";
    import { cn } from "$lib/shadcn";
    import { tick } from "svelte";
    import { fetchDockerImages } from "$lib/index";
    import type { Image, SearchedImage } from "$lib/index";

    export let defaultImages: Image[];
    export let searchedImages: SearchedImage[] = [];
    export let searchTerm: string = "";
    export let value: string = "";
    export let open: boolean = false;

    let timeoutId: ReturnType<typeof setTimeout> | undefined;

    $: selectedDefaultValue =
        defaultImages.find((f: Image) => f.name === value)?.name ??
        "Select an image...";
    $: selectedSearchedValue =
        searchedImages.find((f: SearchedImage) => f.repo_name === value)
            ?.repo_name ?? "Select an image...";

    function closeAndFocusTrigger(triggerId: string) {
        open = false;
        tick().then(() => document.getElementById(triggerId)?.focus());
    }

    const handleInput = (event: Event) => {
        clearTimeout(timeoutId);
        searchTerm = (event.target as HTMLInputElement).value;

        timeoutId = setTimeout(async () => {
            searchedImages = await fetchDockerImages(searchTerm);
        }, 300);
        console.log(searchedImages);
    };

    const handleSelect = (selectedValue: string, triggerId: string) => {
        value = selectedValue;
        closeAndFocusTrigger(triggerId);
    };
</script>

<main>
    <Popover.Root bind:open let:ids>
        <Popover.Trigger asChild let:builder>
            <Button
                builders={[builder]}
                variant="outline"
                role="combobox"
                aria-expanded={open}
                class="w-full justify-between"
            >
                {#if searchTerm === ""}
                    {selectedDefaultValue}
                {:else}
                    {selectedSearchedValue}
                {/if}
                <ChevronsUpDown class="ml-2 h-4 w-4 shrink-0 opacity-50" />
            </Button>
        </Popover.Trigger>
        <Popover.Content class="w-[242px] md:w-[600px] p-0">
            <Command.Root>
                <Input
                    placeholder="Search an Image..."
                    on:input={handleInput}
                />
                <Command.Empty class="border border-gray-800 ">
                    No framework found.
                </Command.Empty>
                <Command.Group class="max-h-[320px] overflow-y-auto ">
                    {#if searchTerm === ""}
                        {#each defaultImages as image}
                            <Command.Item
                                value={image.name}
                                onSelect={() =>
                                    handleSelect(image.name, ids.trigger)}
                            >
                                <Check
                                    class={cn(
                                        "mr-2 h-4 w-4",
                                        value !== image.name &&
                                            "text-transparent",
                                    )}
                                />
                                <div>
                                    <p class="font-bold">{image.name}</p>
                                    <span class="text-xs"
                                        >{image.description}</span
                                    >
                                </div>
                            </Command.Item>
                        {/each}
                    {:else}
                        {#each searchedImages as image}
                            <Command.Item
                                value={image.repo_name}
                                onSelect={() =>
                                    handleSelect(image.repo_name, ids.trigger)}
                            >
                                <Check
                                    class={cn(
                                        "mr-2 h-4 w-4",
                                        value !== image.repo_name &&
                                            "text-transparent",
                                    )}
                                />
                                <div>
                                    <p class="font-bold">{image.repo_name}</p>
                                    <span class="text-xs"
                                        >{image.short_description}</span
                                    >
                                </div>
                            </Command.Item>
                        {/each}
                    {/if}
                </Command.Group>
            </Command.Root>
        </Popover.Content>
    </Popover.Root>
</main>
