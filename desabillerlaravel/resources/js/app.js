/**
 * First we will load all of this project's JavaScript dependencies which
 * includes Vue and other libraries. It is a great starting point when
 * building robust, powerful web applications using Vue and Laravel.
 */

require('./bootstrap');

window.Vue = require('vue');
Vue.config.productionTip = false;

import swal from 'sweetalert2';
window.swal = swal;

import { Form, HasError, AlertError } from 'vform'
Vue.component(HasError.name, HasError)
Vue.component(AlertError.name, AlertError)
window.Form = Form;

window.collect = require('collect.js')
window.moment = require('moment')

import VueApexCharts from 'vue-apexcharts'
window.Apex = {
    noData: {
        text: 'no data...',
    },
    yaxis: {
        labels: {
            minWidth: 40
        }
    },
    tooltip: {
        theme: 'dark',
        x: {
            show: false,
        },
    },
}
Vue.use(VueApexCharts)
Vue.component('apexchart', VueApexCharts)

import VueCurrencyInput from 'vue-currency-input'
Vue.use(VueCurrencyInput, {
    globalOptions: {
        locale: 'de-DE',
        currency: null,
        distractionFree : false
    }
})

Vue.directive('select2', {
    inserted(el) {
        $(el).on('select2:select', () => {
            const event = new Event('change', { bubbles: true, cancelable: true });
            el.dispatchEvent(event);
        });

        $(el).on('select2:unselect', () => {
            const event = new Event('change', {bubbles: true, cancelable: true});
            el.dispatchEvent(event);
        })
    },
});

Vue.directive('daterangepicker', {
    bind(el) {
        $(el).on('apply.daterangepicker', () => {
            const event = new Event('input', {bubbles: true, cancelable: true});
            el.dispatchEvent(event);
        });
    }
});

Vue.directive('anytimepicker', {
    inserted(el) {
        $(el).on('change', () => {
            const event = new Event('input', {bubbles: true, cancelable: true});
            el.dispatchEvent(event);
        });
    }
});
/**
 * The following block of code may be used to automatically register your
 * Vue components. It will recursively scan this directory for the Vue
 * components and automatically register them with their "basename".
 *
 * Eg. ./components/ExampleComponent.vue -> <example-component></example-component>
 */

// const files = require.context('./', true, /\.vue$/i)
// files.keys().map(key => Vue.component(key.split('/').pop().split('.')[0], files(key).default))

// Vue.component('example-component', require('./components/ExampleComponent.vue').default);

/**
 * Next, we will create a fresh Vue application instance and attach it to
 * the page. Then, you may begin adding components to this application
 * or customize the JavaScript scaffolding to fit your unique needs.
 */

// const app = new Vue({
//     el: '#app',
// });
