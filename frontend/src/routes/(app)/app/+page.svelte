<script lang="ts">
    import FormManage from '$lib/components/FormManage/FormManage.svelte';
    import UserGreeting from '$lib/components/UserGreeting.svelte';
    import Statistics from '$lib/components/Statistics.svelte';
</script>

<div class="bg-white min-h-screen py-12 m-6 rounded-2xl px-4">
    <UserGreeting />
    <Statistics />
    <FormManage />   
</div>
