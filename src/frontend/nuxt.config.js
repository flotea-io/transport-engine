import webpack from 'webpack';
export default {
  //mode: 'universal',
  /*
   ** Headers of the page
   */
   env: {
    apiUrl: "https://api.engine.devel"
  },
  head: {
    title: process.env.npm_package_name || '',
    meta: [
    { charset: 'utf-8' },
    { name: 'viewport', content: 'width=device-width, initial-scale=1' },
    {
      hid: 'description',
      name: 'description',
      content: process.env.npm_package_description || ''
    }
    ],
    link: [
    { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' },
    {href:"https://fonts.googleapis.com/css?family=Open+Sans:300,400,600,700,800", rel:"stylesheet"}
    ]
  },
  /*
   ** Customize the progress-bar color
   */
   loading: { color: '#fff' },
  /*
   ** Global CSS
   */
   css: [
   'element-ui/lib/theme-chalk/index.css',
   'jgrowl/jquery.jgrowl.css',
   'codemirror/lib/codemirror.css',
   'codemirror/addon/merge/merge.css',
   'codemirror/theme/eclipse.css'
   ],
  /*
   ** Plugins to load before mounting the App
   */
   plugins: [
   '@/plugins/element-ui',
   { src: '~plugins/vue-full-calendar', ssr: false },
   { src: '~plugins/nuxt-codemirror-plugin.js', ssr: false }
   ],
  /*
   ** Nuxt.js dev-modules
   */
   buildModules: [
    // Doc: https://github.com/nuxt-community/eslint-module
    //'@nuxtjs/eslint-module'
    ],
  /*
   ** Nuxt.js modules
   */
   modules: ['nuxt-leaflet'],
  /*
   ** Build configuration
   */
   build: {
    //extractCSS: true,
    plugins: [
    new webpack.ProvidePlugin({  
      $: 'jquery',
      jQuery: 'jquery',
    }),


    ],
    transpile: [/^element-ui/],
    /*
     ** You can extend webpack config here
     */
     extend(config, ctx) {

     }
   }
 }
