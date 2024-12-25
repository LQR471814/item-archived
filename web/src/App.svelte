<script lang="ts">
  import Fs from "./lib/FS.svelte";
  import { archive } from "./archive";
</script>

<main>
  <h1 class="font-bold text-2xl">Item Archive</h1>
  <Fs {archive} />
</main>
