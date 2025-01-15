/**
 * @file vite.config.js
 *
 * @copyright © 2024 s3mat3
 * This code is licensed under the MIT License, see the LICENSE file for details
 *
 * @brief
 *
 * @author s3mat3
 */
import { defineConfig } from 'vite';
import { resolve } from 'path';

const root       = resolve(__dirname, 'src');
const publicDir  = resolve(__dirname, 'public');

export default defineConfig({
    root,
    base: './',
    publicDir: true,
    server: {
        host: '0.0.0.0',
        port: 10080
    },
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
