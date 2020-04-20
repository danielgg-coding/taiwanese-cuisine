import React from 'react'

function CenterMessage({text}) {
  return (
    <div style={{ paddingTop: '6%', paddingBottom:  '4%',  textAlign: 'center', fontSize: 20 }}>
     {text}
    </div>
  )
}

export default CenterMessage