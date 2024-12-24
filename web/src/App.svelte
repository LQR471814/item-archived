<script lang="ts">
  import Fs from "./lib/FS.svelte";
  import { archive } from "./archive";
</script>

<main>
  <Fs {archive} />
</main>
