<template>
  <div class="relative w-screen h-screen overflow-hidden">
    <div class="absolute inset-0 -z-10">
      <svg aria-hidden="true"
        class="w-full h-full stroke-stone-200 [mask-image:radial-gradient(100%_100%_at_top,white,transparent)]">
        <defs>
          <pattern id="grid-pattern" x="50%" y="-1" width="200" height="200" patternUnits="userSpaceOnUse">
            <path d="M100 200V0.5M0.5 0.5H200" fill="none" />
          </pattern>
        </defs>

        <rect width="100%" height="100%" fill="url(#grid-pattern)" stroke-width="0" />
      </svg>
    </div>

    <div class="relative z-10">
      <slot />
    </div>
  </div>
</template>
