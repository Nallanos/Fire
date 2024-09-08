<script>
    import { Button } from "$lib/components/ui/button";
</script>

<nav class="h-24 items-center flex w-full justify-between container mx-auto">
    <a href="/">
        <img src="/Fire.png" alt="Fire logo" />
    </a>
    <div class="flex items-center gap-4 pr-6">
        <a href="/apps">
            <button>Apps</button>
        </a>
        <form action="?/deleteCookies" method="post">
            <Button type="submit" variant="outline">Logout</Button>
        </form>
    </div>
</nav>

<slot />
