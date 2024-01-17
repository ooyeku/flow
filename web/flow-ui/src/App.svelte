<script lang="ts">
  // Import your components here
  import Tasks from './components/Tasks.svelte';
  // import Planning from './Planning.svelte';
  import Modal from "./components/Modal.svelte";

  let showModal = false;
</script>

<style>
  /* Define your styles here */
  .app {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    height: 100vh;
    font-family: Arial, sans-serif;
    background-color: #f5f5f5;
  }

  .header {
    font-size: 2em;
    color: #333;
  }

  .section {
    width: 80%;
    padding: 1em;
    margin: 1em 0;
    background-color: #fff;
    border-radius: 8px;
    box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
  }

  .section h2 {
    margin: 0 0 1em;
    color: #666;
  }

  /* Add more styles as needed */
</style>

<div class="app">
  <div class="header">Welcome to Flow</div>
  <div class="section">
    <h2>Tasks</h2>
<!--     <Tasks />-->
  </div>
  <div class="section">
    <h2>Planning</h2>
    <!-- <Planning /> Uncomment this when you have implemented Planning component -->
    <p>Planning will go here</p>
  </div>
</div>