// @ts-check
// Note: type annotations allow type checking and IDEs autocompletion
const path = require('path');
const jsdom = require('jsdom');

const lightCodeTheme = require('prism-react-renderer/themes/github');
const darkCodeTheme = require('prism-react-renderer/themes/dracula');

// Emulate DOMParser with jsdom
const { JSDOM } = jsdom;
global.DOMParser = new JSDOM().window.DOMParser;

/** @type {import('@docusaurus/types').Config} */
const config = {
	title: 'Sandbox docs',
	tagline: 'POSH Sandbox',
	favicon: 'img/favicon.ico',
	url: 'https://localhost:3000',
	baseUrl: '/',

	organizationName: 'foomo',
	projectName: 'posh-sandbox',

	onBrokenLinks: 'warn',
	onBrokenMarkdownLinks: 'warn',

	trailingSlash: false,

	i18n: {
		defaultLocale: 'en',
		locales: ['en'],
	},

	presets: [
		[
			'classic',
			/** @type {import('@docusaurus/preset-classic').Options} */
			({
				blog: false,
				pages: false,
				docs: {
					path: 'docs',
					routeBasePath: '/',
					// Please change this to your repo.
					// Remove this to remove the "edit this page" links.
					editUrl: undefined,
					include: ['**/*.md'],
					exclude: ['**/.*/**', '**/bin/**', '**/tmp/**', '**/node_modules/**'],
				},
				theme: {
					customCss: require.resolve('./src/css/custom.css'),
				},
			}),
		],
	],
	themes: ['@docusaurus/theme-live-codeblock'],
	themeConfig:
		/** @type {import('@docusaurus/preset-classic').ThemeConfig} */
		({
			navbar: {
				title: 'Sandbox',
				items: [
					{
						href: 'https://github.com/foomo/posh-sandbox',
						label: 'GitHub',
						position: 'right',
					},
				],
			},
			prism: {
				theme: lightCodeTheme,
				darkTheme: darkCodeTheme,
			},
		}),
};

module.exports = config;
