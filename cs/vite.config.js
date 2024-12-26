import { defineConfig } from 'vite';
import { resolve } from 'path';

const root       = resolve(__dirname, 'src');
const publicDir  = resolve(__dirname, 'public');

export default defineConfig({
    root,
    base: './',
    publicDir: true,
    resolve: {
        alias: {
            '@pages' : resolve(root, 'pages'),
            '@assets': resolve(root, 'assets'),
            '@public': publicDir,
            '@lib'   : resolve(root, 'lib'),
            '@comp'  : resolve(root, 'components'),
        },
    },
    css: {
        preprocessorOptions: {
            scss: {
                api: "modern-compiler",
                // additionalData: `@import "@assets/styles/srl.scss";`,
            },
        },
    },
    build: {
        outDir: '../dist',
        emptyOutDir: true,
        rollupOptions: {
            input: {
                // manage: resolve(root, 'pages/manage', 'index.html'),
                showcase: resolve(root, 'pages/showcase', 'index.html'),
                // example: resolve(root, 'pages/example', 'index.html'),
            },
            output: {
                entryFileNames: `assets/[name]/entrypoint.js`,
                chunkFileNames: `assets/[name]/chank.js`,
                assetFileNames: `assets/[name]/asset.[ext]`,
            },
        },
    },
    test: {
        globals: true,
        environment: 'jsdom',
        include: ['./**/tests/*.test.js']
    },
})
