// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import BelshareBelshare from './belshare.belshare'
import BelshareEav from './belshare.eav'
import BelshareUser from './belshare.user'


export default { 
  BelshareBelshare: load(BelshareBelshare, 'belshare.belshare'),
  BelshareEav: load(BelshareEav, 'belshare.eav'),
  BelshareUser: load(BelshareUser, 'belshare.user'),
  
}


function load(mod, fullns) {
    return function init(store) {        
        if (store.hasModule([fullns])) {
            throw new Error('Duplicate module name detected: '+ fullns)
        }else{
            store.registerModule([fullns], mod)
            store.subscribe((mutation) => {
                if (mutation.type == 'common/env/INITIALIZE_WS_COMPLETE') {
                    store.dispatch(fullns+ '/init', null, {
                        root: true
                    })
                }
            })
        }
    }
}