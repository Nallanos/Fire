<script lang="ts">
    import { Button } from "$lib/components/ui/button/index.js";
    import * as Card from "$lib/components/ui/card/index.js";
    import * as Select from "$lib/components/ui/select/index.js";
    import { Input } from "$lib/components/ui/input/index.js";
    import { Label } from "$lib/components/ui/label/index.js";
    import { createApp } from "$lib/api/apps/createApp";
    let appName = "";
    async function updateInputVal(event: Event) {
        let inputElement = event.target as HTMLInputElement;
        appName = inputElement.value;
    }
    const frameworks = [
        {
            value: "sveltekit",
            label: "SvelteKit",
        },
        {
            value: "next",
            label: "Next.js",
        },
        {
            value: "astro",
            label: "Astro",
        },
        {
            value: "nuxt",
            label: "Nuxt.js",
        },
    ];
</script>

<main class="flex container justify-center pt-12 h-full">
    <Card.Root class="w-[750px] border border-gray-800 ">
        <Card.Header>
            <Card.Title>Create project</Card.Title>
            <Card.Description
                >Deploy your new project in one-click.</Card.Description
            >
        </Card.Header>
        <Card.Content>
            <form
                on:submit={async () => {
                    await createApp({ name: appName });
                    window.location.href = "/apps";
                }}
            >
                <div class="grid w-full items-center gap-4">
                    <div class="flex flex-col space-y-1.5">
                        <Label for="name">Name</Label>
                        <Input
                            on:input={updateInputVal}
                            bind:value={appName}
                            id="name"
                            placeholder="Name of your project"
                            class="focus:outline-none"
                        />
                    </div>
                    <div class="flex flex-col space-y-1.5">
                        <Label for="framework">Framework</Label>
                        <Select.Root>
                            <Select.Trigger id="framework">
                                <Select.Value placeholder="Select" />
                            </Select.Trigger>
                            <Select.Content>
                                {#each frameworks as framework}
                                    <Select.Item
                                        value={framework.value}
                                        label={framework.label}
                                        >{framework.label}</Select.Item
                                    >
                                {/each}
                            </Select.Content>
                        </Select.Root>
                    </div>
                    <Card.Footer class="flex justify-between p-0">
                        <Button variant="outline">Cancel</Button>
                        <Button type="submit">Create</Button>
                    </Card.Footer>
                </div>
            </form>
        </Card.Content>
    </Card.Root>
</main>
