<script setup>
import { ref } from "vue";


const golyInput = ref("") 

async function onSubmit(event) {
  console.log(golyInput.value);
  const rawResponse = await fetch('http://localhost:3030/goly', {
    method: 'POST',
    headers: {
      'Accept': 'application/json',
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({redirect:golyInput.value})
  });
  const content = await rawResponse.json();
  console.log(content)
}

</script>

<template>

    <header class="relative z-50 w-full h-24">
        <div class="container flex items-center justify-center h-full  px-8 mx-auto sm:justify-between xl:px-0">

            <a href="/" class="relative flex items-center inline-block h-5 h-full font-black leading-none">
                <img src="./assets/logo.png">
            </a>

            <svg class=" hidden w-screen max-w-2xl -mt-72 -ml-12 lg:block" viewBox="0 0 818 815"
                xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink">
                <defs>
                    <linearGradient x1="0%" y1="0%" x2="100%" y2="100%" id="c">
                        <stop stop-color="#E614F2" offset="0%" />
                        <stop stop-color="#FC3832" offset="100%" />
                    </linearGradient>
                    <linearGradient x1="0%" y1="0%" x2="100%" y2="100%" id="f">
                        <stop stop-color="#657DE9" offset="0%" />
                        <stop stop-color="#1C0FD7" offset="100%" />
                    </linearGradient>
                    <filter x="-4.7%" y="-3.3%" width="109.3%" height="109.3%" filterUnits="objectBoundingBox" id="a">
                        <feOffset dy="8" in="SourceAlpha" result="shadowOffsetOuter1" />
                        <feGaussianBlur stdDeviation="8" in="shadowOffsetOuter1" result="shadowBlurOuter1" />
                        <feColorMatrix values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.15 0" in="shadowBlurOuter1" />
                    </filter>
                    <filter x="-4.7%" y="-3.3%" width="109.3%" height="109.3%" filterUnits="objectBoundingBox" id="d">
                        <feOffset dy="8" in="SourceAlpha" result="shadowOffsetOuter1" />
                        <feGaussianBlur stdDeviation="8" in="shadowOffsetOuter1" result="shadowBlurOuter1" />
                        <feColorMatrix values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0.2 0" in="shadowBlurOuter1" />
                    </filter>
                    <path
                        d="M160.52 108.243h497.445c17.83 0 24.296 1.856 30.814 5.342 6.519 3.486 11.635 8.602 15.12 15.12 3.487 6.52 5.344 12.985 5.344 30.815v497.445c0 17.83-1.857 24.296-5.343 30.814-3.486 6.519-8.602 11.635-15.12 15.12-6.52 3.487-12.985 5.344-30.815 5.344H160.52c-17.83 0-24.296-1.857-30.814-5.343-6.519-3.486-11.635-8.602-15.12-15.12-3.487-6.52-5.343-12.985-5.343-30.815V159.52c0-17.83 1.856-24.296 5.342-30.814 3.486-6.519 8.602-11.635 15.12-15.12 6.52-3.487 12.985-5.343 30.815-5.343z"
                        id="b" />
                    <path
                        d="M159.107 107.829H656.55c17.83 0 24.296 1.856 30.815 5.342 6.518 3.487 11.634 8.602 15.12 15.12 3.486 6.52 5.343 12.985 5.343 30.816V656.55c0 17.83-1.857 24.296-5.343 30.815-3.486 6.518-8.602 11.634-15.12 15.12-6.519 3.486-12.985 5.343-30.815 5.343H159.107c-17.83 0-24.297-1.857-30.815-5.343-6.519-3.486-11.634-8.602-15.12-15.12-3.487-6.519-5.343-12.985-5.343-30.815V159.107c0-17.83 1.856-24.297 5.342-30.815 3.487-6.519 8.602-11.634 15.12-15.12 6.52-3.487 12.985-5.343 30.816-5.343z"
                        id="e" />
                </defs>
                <g fill="none" fill-rule="evenodd" opacity=".9">
                    <g transform="rotate(65 416.452 409.167)">
                        <use fill="#000" filter="url(#a)" xlink:href="#b" />
                        <use fill="url(#c)" xlink:href="#b" />
                    </g>
                    <g transform="rotate(29 421.929 414.496)">
                        <use fill="#000" filter="url(#d)" xlink:href="#e" />
                        <use fill="url(#f)" xlink:href="#e" />
                    </g>
                </g>
            </svg>




        </div>
    </header>
    <!-- End Header Section-->
    <div class="relative items-center justify-center w-full overflow-x-hidden lg:pt-20 lg:pb-40 xl:pt-16 xl:pb-64">
        <div class="container flex flex-col items-center justify-between h-full max-w-6xl px-8 mx-auto -mt-32 lg:flex-row xl:px-0">
            <div class="z-30 flex flex-col items-center w-full max-w-xl pt-48 text-center lg:items-start  lg:pt-20 xl:pt-40 lg:text-left">
                <h1 class="relative mb-4 text-3xl font-black leading-tight text-gray-900 sm:text-6xl xl:mb-8">
                    Short URLS`s are Great Thing</h1>
                <div>
                    <form @submit.prevent="onSubmit" class="m-4 flex">
                        <input v-model="golyInput"  class="rounded-l-lg p-4 border-t mr-0 border-b border-l text-gray-800 border-gray-200 bg-white"
                            placeholder="Long URL" />
                        <button
                            class="px-8 rounded-r-lg bg-indigo-600  text-base font-bold text-white  font-bold p-4 uppercase  border-t border-b border-r">Shorten</button>
                    </form>
                </div>
                <!-- <a href="#_"
                    class="relative self-start inline-block w-auto px-8 py-4 mx-auto mt-0 text-base font-bold text-white bg-indigo-600 border-t border-gray-200 rounded-md shadow-xl sm:mt-1 fold-bold lg:mx-0">Signup
                    Today!</a> -->
                <!-- Integrates with section -->

            </div>
            
        </div>
    </div>
</template>

