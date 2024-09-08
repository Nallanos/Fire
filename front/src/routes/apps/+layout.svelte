<script>
    import { Button } from "$lib/components/ui/button";
</script>

<header class="container flex items-center mx-auto px-4">
    <nav class="h-24 items-center flex w-full justify-between">
        <a href="/" class="pl-36">
            <img src="/Fire.png" alt="Fire logo" />
        </a>
        <div class="flex items-center gap-4">
            <a href="/">
                <button>Home page</button>
            </a>
            <a href="/apps">
                <button>Apps</button>
            </a>
            <form action="?/deleteCookies" method="post">
                <Button type="submit" variant="outline">Logout</Button>
            </form>
        </div>
    </nav>
</header>

<slot />
