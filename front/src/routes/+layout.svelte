<script>
  import "../app.css";

  const items = [
    { name: "Home", href: "/" },
    { name: "Services", href: "/services" },
    { name: "Apps", href: "/apps" },
  ];
</script>

<main class="text-white bg-gray-950 min-h-screen">
  <header class="border-b border-gray-800">
    <div
      class="flex items-center container mx-auto lg:mx-auto lg:justify-between md: mx-auto md:justify-between min-[450px]:justify-between min-[450px]:mx-auto pr-4"
    >
      <a href="/">
        <img src="/Fire.png" alt="Fire logo" class="size-24" />
      </a>

      <ul class="flex lg:gap-20 gap-8 text-md">
        {#each items as item}
          <li>
            <a class="hover:text-gray-300" href={item.href}>
              {item.name}
            </a>
          </li>
        {/each}
      </ul>
    </div>
  </header>

  <div class="container mx-auto px-4">
    <slot />
  </div>
</main>
