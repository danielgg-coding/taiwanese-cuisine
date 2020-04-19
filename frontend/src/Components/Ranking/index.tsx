import React, { useMemo, useState } from 'react';
import {
  FacebookShareButton,
  LineShareButton,
  TwitterShareButton,
  FacebookIcon,
  LineIcon,
  TwitterIcon,
} from "react-share";
import { DataView } from '@aragon/ui'
import CroppedImg from '../common/CroppedImg'

import { getList } from '../../api';
import type { food } from '../../types'

function Ranking() {

  const [foodList, setList] = useState<food[]>([])

  useMemo(async () => {
    const _list = await getList()
    const sortedList = _list.sort((a, b) => a.score > b.score ? -1 : 1)
    setList(sortedList)
  }, [])

  return (
    <>
      <div style={{ padding: '6%', textAlign: 'center', fontSize: 36 }}>
        小吃排行榜
      </div>
      <DataView
        fields={['排名', '', '名稱', '分數', '出賽']}
        entries={foodList.slice(0, 10)}
        renderEntry={(item: food, index: number) => {
          return [
            index + 1,
            <div style={{ paddingTop: '6%', paddingBottom: '4%', textAlign: 'center', fontSize: 25 }}>
              {item.name}
            </div>,
            <CroppedImg url={item.image} width={250} height={100} />,
            item.score,
            item.played
          ]
        }}
        entriesPerPage={5}
      />
      <div style={{ paddingTop: '6%', paddingBottom: '4%', textAlign: 'center', fontSize: 20 }}>
        <FacebookShareButton
          url="https://github.com/nygardk/react-share/blob/master/demo/Demo.tsx"
          quote="台灣小吃爭霸戰"
        >
          <FacebookIcon size={50} round />
        </FacebookShareButton>
        <LineShareButton
          url="https://github.com/nygardk/react-share/blob/master/demo/Demo.tsx"
          title="萬人投票結果，台灣小吃滷肉飯只排第二名...?！"
        >
          <LineIcon size={50} round />
        </LineShareButton>

        <TwitterShareButton
          url="https://github.com/nygardk/react-share/blob/master/demo/Demo.tsx"
          title="萬人投票結果，台灣小吃滷肉飯只排第二名...?！"
        >
          <TwitterIcon size={50} round />
        </TwitterShareButton>
      </div>
    </>
  );
}

export default Ranking;
