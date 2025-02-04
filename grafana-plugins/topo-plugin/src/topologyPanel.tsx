import React, { useRef, useState, useEffect } from 'react';
import * as G6 from '@antv/g6';
import _ from 'lodash';
import './topology/node';
import './topology/edge';
import { formatTime, formatCount, formatKMBT, formatPercent, nodeTooltip } from './topology/tooltip';
import TopoLegend from './topology/legend';
import { metricList, layoutOptions, directionOptions, viewRadioOptions, showServiceOptions, NodeDataProps, EdgeDataProps, 
  buildLayout, transformData, nsRelationHandle, workloadRelationHandle } from './topology/services'; 
import { PanelProps } from '@grafana/data';
import { SimpleOptions } from 'types';
import { css, cx } from 'emotion';
import { stylesFactory, Select, RadioButtonGroup, Icon, Tooltip, Spinner } from '@grafana/ui';

interface VolumeProps {
  maxSentVolume: number; 
  maxReceiveVolume: number;
  minSentVolume: number; 
  minReceiveVolume: number;
}

let SGraph: any;
let topoData: any, nodeData: NodeDataProps, edgeData: EdgeDataProps;
interface Props extends PanelProps<SimpleOptions> {}
export const TopologyPanel: React.FC<Props> = ({ options, data, width, height, replaceVariables }) => {
  let graphRef: any = useRef();
  // const theme = useTheme();
  const namespace = replaceVariables('$namespace');
  const workload = replaceVariables('$workload');
  const styles = getStyles();
  const [graphData, setGraphData] = useState<any>({}); 
  const [layout, setLayout] = useState<string>('dagre'); 
  const [loading, setLoading] = useState<boolean>(true); 
  const [showCheckbox, setShowCheckbox] = useState<boolean>(namespace.split(',').length === 1);
  const [showService, setShowService] = useState<boolean>(false);
  const [showView, setShowView] = useState<boolean>(false);
  const [firstChangeDir, setFirstChangeDir] = useState<boolean>(false);
  const [direction, setDirection] = useState<string>('LR');
  const [view, setView] = useState<string>('workload_view');
  const [lineMetric, setLineMetric] = useState<any>('latency');
  const [volumes, setVolumes] = useState<VolumeProps>({maxSentVolume: 0, maxReceiveVolume: 0, minSentVolume: 0, minReceiveVolume: 0});
  const [nodeTypesList, setNodeTypesList] = useState<any[]>([]);

  // console.log(options, namespace, workload, width, height);
  // console.log(data);

  // 当勾选View Service Call时，显示service的调用边，两个节点之间存在多条调用关系，使用弧线绘制对应的调用关系
  const serviceLineUpdate = (dir = direction) => {
    let activeList: any[] = [];
    const edges = SGraph.getEdges();
    const offest = 5;
    edges.forEach((edge: any) => {
      let edgeModel = edge.getModel();
      let active = activeList.findIndex((item: any) => (item.source === edgeModel.source && item.target === edgeModel.target) || (item.source === edgeModel.target && item.target === edgeModel.source));
      if (active === -1) {
        activeList.push({
          source: edgeModel.source,
          target: edgeModel.target
        });
        let lines = edges.filter((itemEdge: any) => {
          let item = itemEdge.getModel();
          return (item.source === edgeModel.source && item.target === edgeModel.target) || (item.source === edgeModel.target && item.target === edgeModel.source)
        });
        if (lines.length > 1) {
          let oddNum = 0, evenNum = 0;
          lines.forEach((item: any, idx: number) => {
            let line: any = item.getContainer();
            // line.type = 'service-edge2';
            let curveOffset = 0;
            // if (idx % 2 === 0) {
            //   curveOffset = arc * (1 + (1 * evenNum));
            //   evenNum ++;
            // } else {
            //   curveOffset = -arc * (1 + (1 * oddNum));
            //   oddNum ++;
            // }
            // console.log(item, curveOffset)
            // line.curveOffset = curveOffset;
            // SGraph.updateItem(item, line);
            if (idx % 2 === 0) {
              curveOffset = -offest * (1 + (1 * evenNum));
              evenNum ++;
            } else {
              curveOffset = offest * (1 + (1 * oddNum));
              oddNum ++;
            }
            if (dir === 'TB') {
              line.translate(curveOffset, 0);
            } else {
              line.translate(0, curveOffset);
            }
          });
        }
      }
    });
  }
  /**
   * update edge style based on the current metric
   * 根据当前指标选择更新边的样式
   */
  const updateLinesAndNodes = (metric = lineMetric, serviceLine = showService) => {
    const nodes = SGraph.getNodes();
    const edges = SGraph.getEdges();
    if (metric === 'latency' || metric === 'rtt' || metric === 'errorRate') {
      edges.forEach((edge: any, idx: number) => {
        let edgeModel = edge.getModel();
        let color: string;
        
        if (metric === 'latency') {
          color = edgeModel.latency > options.abnormalLatency ? '#ff4c4c' : (edgeModel.latency > options.normalLatency ? '#f3ff69' : '#C2C8D5');
          edgeModel.label = formatTime(edgeModel.latency);
        } else if (metric === 'rtt') {
          color = edgeModel.rtt > options.abnormalRtt ? '#ff4c4c' : (edgeModel.rtt > options.normalRtt ? '#f3ff69' : '#C2C8D5');
          edgeModel.label = formatTime(edgeModel.rtt);
        } else {
          color = edgeModel.errorRate > 0 ? '#ff4c4c' : '#C2C8D5';
          edgeModel.label = formatPercent(edgeModel.errorRate);
        }
        edgeModel.opposite && (edgeModel.labelCfg.refY = -10);
        edgeModel.style.stroke = color;
        if (serviceLine) {
          edgeModel.rectColor = color;
        }
        edgeModel.style.lineWidth = 1;
        SGraph.updateItem(edge, edgeModel);
      });
      nodes.forEach((node: any) => {
        let nodeModel = node.getModel();
        if (metric === 'latency') {
          nodeModel.status = nodeModel.latency > options.abnormalLatency ? 'red' : (nodeModel.latency > options.normalLatency ? 'yellow' : 'green');
        } else if (metric === 'rtt') {
          nodeModel.status = 'green';
        } else {
          nodeModel.status = nodeModel.errorRate > 0 ? 'red' : 'green';
        }
        SGraph.updateItem(node, nodeModel);
      });
    } else if (metric === 'sentVolume' || metric === 'receiveVolume'){
      if (metric === 'sentVolume') {
        let volumeStep = volumes.maxSentVolume / 5;
        if (edges.length === 1) {
          let edge = edges[0];
          let edgeModel = edge.getModel();
          edgeModel.style.lineWidth = 1;
          edgeModel.style.stroke = '#C2C8D5';
          edgeModel.label = formatKMBT(edgeModel.sentVolume);
          SGraph.updateItem(edge, edgeModel);
        } else {
          edges.forEach((edge: any, idx: number) => {
            let edgeModel = edge.getModel();
            let step = Math.floor(edgeModel.sentVolume / volumeStep);
            edgeModel.style.lineWidth = step === 0 ? 1 : 1.5 * step;
            edgeModel.style.stroke = '#C2C8D5';
            edgeModel.label = formatKMBT(edgeModel.sentVolume);
            edgeModel.opposite && (edgeModel.labelCfg.refY = -10);
            SGraph.updateItem(edge, edgeModel);
          });
        }
      } else {
        let volumeStep = volumes.maxReceiveVolume / 5;
        if (edges.length === 1) {
          let edge = edges[0];
          let edgeModel = edge.getModel();
          edgeModel.style.lineWidth = 1;
          edgeModel.style.stroke = '#C2C8D5';
          edgeModel.label = formatKMBT(edgeModel.receiveVolume);
          SGraph.updateItem(edge, edgeModel);
        } else {
          edges.forEach((edge: any, idx: number) => {
            let edgeModel = edge.getModel();
            let step = Math.floor(edgeModel.receiveVolume / volumeStep);
            edgeModel.style.lineWidth = step === 0 ? 1 : 1.5 * step;
            edgeModel.style.stroke = '#C2C8D5';
            edgeModel.label = formatKMBT(edgeModel.receiveVolume);
            edgeModel.opposite && (edgeModel.labelCfg.refY = -10);
            SGraph.updateItem(edge, edgeModel);
          });
        }
      }
      nodes.forEach((node: any) => {
        let nodeModel = node.getModel();
        nodeModel.status = 'green';
        SGraph.updateItem(node, nodeModel);
      });
    } else {
      edges.forEach((edge: any) => {
        let edgeModel = edge.getModel();
        edgeModel.style.stroke = '#C2C8D5';
        edgeModel.style.lineWidth = 1;
        edgeModel.label = formatCount(edgeModel[metric]);
        SGraph.updateItem(edge, edgeModel);
      });
      nodes.forEach((node: any) => {
        let nodeModel = node.getModel();
        nodeModel.status = 'green';
        SGraph.updateItem(node, nodeModel);
      });
    }
  }
  // draw topology
  const draw = (gdata: any, serviceLine = showService) => {
    const inner: any = document.getElementById('kindling_topo');
    inner.innerHTML = '';
    let data = _.cloneDeep(gdata);
    const graph = new G6.Graph({
      // renderer: 'svg',
      container: 'kindling_topo',
      width: width - 240,
      height: height,
      fitView: true,
      fitViewPadding: 10,
      maxZoom: 1.5,
      minZoom: 0.25,
      fitCenter: true,
      autoPaint: false,
      plugins: [nodeTooltip],
      modes: {
        default: [
          {
            type: 'drag-canvas',
            enableOptimize: true,
          }, {
            type: 'zoom-canvas',
            maxZoom: 1.5,
            minZoom: 0.25
          }, 
          'drag-node'
        ]
      },
      layout: buildLayout(layout, direction),
      defaultNode: {
        type: 'custom-node'
      },
      defaultEdge: {
        type: 'service-edge',
        labelCfg: {
          refY: serviceLine ? 15 : 10,
          autoRotate: true,
          style: {
            fontWeight: 400,
            fill: '#C2C8D5',
          }
        },
        style: {
          radius: 10,
          offset: 5,
          endArrow: true,
          lineWidth: 1,
          stroke: '#C2C8D5',
        }
      }
    });
    graph.data(data);
    graph.render();

    SGraph = graph;
    serviceLineUpdate();
    updateLinesAndNodes(lineMetric, serviceLine);
  };
 
  /**
   * Gets an array of node types in the current topology for legend drawing on the right
   * 获取当前拓扑图下节点的类型数组，用于右侧的legend绘制
   */
  const getNodeTypes = (nodes: any[]) => {
    let nodeByType = _.groupBy(nodes, 'nodeType');
    let types: string[] = _.keys(nodeByType);
    return types;
  }
  /**
   * When you go back to the topology drawing data, update the type array of the corresponding node and the value of Max and min of the flow on the side
   * 重新回去拓扑绘制数据时，更新对应的节点的类型数组和边上流量max、min的数值
   */
  const handleResult = (gdata: any) => {
    let nodeTypesList = getNodeTypes(gdata.nodes);
    setNodeTypesList(nodeTypesList);
    let volumeData: VolumeProps = {
      maxSentVolume: _.max(_.map(gdata.edges, 'sentVolume')),
      maxReceiveVolume: _.max(_.map(gdata.edges, 'receiveVolume')),
      minSentVolume: _.min(_.map(gdata.edges, 'sentVolume')),
      minReceiveVolume: _.min(_.map(gdata.edges, 'receiveVolume'))
    }
    setVolumes(volumeData);
  }
  // Initial data processing: Generates topology data
  const initData = () => {
    let nodes: any[] = [], edges: any[] = [];
    topoData = transformData(_.filter(data.series, (item: any) => item.refId === 'A'));
    let edgeTimeData: any = transformData(_.filter(data.series, (item: any) => item.refId === 'I'));
    let edgeSendVolumeData: any = transformData(_.filter(data.series, (item: any) => item.refId === 'B'));
    let edgeReceiveVolumeData: any = transformData(_.filter(data.series, (item: any) => item.refId === 'C'));
    let edgeRetransmitData: any = transformData(_.filter(data.series, (item: any) => item.refId === 'J'));
    let edgeRTTData: any = transformData(_.filter(data.series, (item: any) => item.refId === 'K'));
    let edgePackageLostData = transformData(_.filter(data.series, (item: any) => item.refId === 'F'));
    edgeData = {
      edgeCallData: topoData,
      edgeTimeData,
      edgeSendVolumeData,
      edgeReceiveVolumeData,
      edgeRetransmitData,
      edgeRTTData,
      edgePackageLostData
    };
    
    let nodeCallsData: any = transformData(_.filter(data.series, (item: any) => item.refId === 'D'));
    let nodeTimeData: any = transformData(_.filter(data.series, (item: any) => item.refId === 'E'));
    let nodeSendVolumeData: any = transformData(_.filter(data.series, (item: any) => item.refId === 'G'));
    let nodeReceiveVolumeData: any = transformData(_.filter(data.series, (item: any) => item.refId === 'H'));
    nodeData = {
      nodeCallsData,
      nodeTimeData,
      nodeSendVolumeData,
      nodeReceiveVolumeData
    };
    console.log('edgeData', edgeData);
    console.log('nodeData', nodeData);
    if (namespace.indexOf(',') > -1) {
      let result: any = nsRelationHandle(topoData, nodeData, edgeData);
      nodes = [].concat(result.nodes);
      edges = [].concat(result.edges);
    } else {
      let showPod = workload.split(',').length === 1;
      setView(showPod ? 'pod_view' : 'workload_view');
      let result: any = workloadRelationHandle(workload, namespace, topoData, nodeData, edgeData, showPod, showService);
      nodes = [].concat(result.nodes);
      edges = [].concat(result.edges);
    }
    
    let gdata = {
      nodes: nodes,
      edges: edges
    }
    console.log(gdata);
    setGraphData(gdata);
    draw(gdata);
    handleResult(gdata);
  }

  useEffect(() => {
    if (SGraph) {
      draw(graphData);
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [layout]);
  useEffect(() => {
    if (SGraph) {
      updateLinesAndNodes();
    }
  // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [volumes]);

  useEffect(() => {
    setLoading(true);
    if (namespace.split(',').length === 1) {
      setShowCheckbox(true);
      if (workload.split(',').length === 1) {
        setShowView(true);
      } else {
        setShowView(false);
      }
    } else {
      setShowCheckbox(false);
      setShowView(false);
    }
  }, [namespace, workload]);
  useEffect(() => {
    if (data.state === 'Done') {
      setLoading(false);
      initData();
    }
	// eslint-disable-next-line react-hooks/exhaustive-deps
  }, [data, namespace]);

  // When the segment indicator is switched, the corresponding segment style is updated
  const lineMetricChange = (opt: any) => {
    setLineMetric(opt.value);
    updateLinesAndNodes(opt.value);
  }
  const changeLayout = (opt: any) => {
    if (opt.value === layout) {
      return;
    }
    setLayout(opt.value);
    // draw(graphData);
  } 
  const changeDirection = (value: any) => {
    setDirection(value);
    SGraph.updateLayout({
      rankdir: value
    });
    SGraph.fitView(10);
    if (!firstChangeDir) {
      serviceLineUpdate(value);
      setFirstChangeDir(true);
    }
  }
  // Whether to display service calls on invocation relationships
  const changeShowService = () => {
    let show = !showService ? true : false;
    setShowService(show);
    let { nodes, edges } = workloadRelationHandle(workload, namespace, topoData, nodeData, edgeData, view === 'pod_view', show);
    let gdata = {
      nodes: nodes,
      edges: edges
    }
    draw(gdata, show);
    setGraphData(gdata);
    handleResult(gdata);
  }
  // toggle View Mode。Switch between the workload view and pod view
  const changeView = (value: any) => {
    setView(value);
    let { nodes, edges } = workloadRelationHandle(workload, namespace, topoData, nodeData, edgeData, value === 'pod_view', showService);
    let gdata = {
      nodes: nodes,
      edges: edges
    }
    draw(gdata);
    setGraphData(gdata);
    handleResult(gdata);
  }

  return (
    <div
      className={cx(
        styles.wrapper,
        css`
          width: ${width}px;
          height: ${height}px;
        `
      )}
    >
      <div className={styles.topLineMetric}>
        <div className={styles.metricSelect}>
          <span style={{ width: '180px' }}>Call Line Metric</span>
          <Select value={lineMetric} options={metricList} onChange={lineMetricChange}/>
        </div>
      </div>
      <div className={styles.topRightWarp}>
        <div className={styles.viewRadioMode}>
          <div>
            <span>Layout</span>
            <Tooltip content="change topology layout。">
              <Icon name="question-circle" />
            </Tooltip>
          </div>
          <div style={{ width: 150 }}>
            <Select value={layout} options={layoutOptions} onChange={changeLayout}/>
          </div>
        </div>
        {
          layout === 'dagre' ? <div className={styles.viewRadioMode}>
            <div>
              <span>Layout Direction</span>
              <Tooltip content="change Dargre topology layout direction mean top to bottom，LR mean left to right。">
                <Icon name="question-circle" />
              </Tooltip>
            </div>
            <RadioButtonGroup options={directionOptions} value={direction} onChange={changeDirection}/>
          </div> : null
        }
        {
          showView ? <div className={styles.viewRadioMode}>
            <span>View Mode</span>
            <RadioButtonGroup options={viewRadioOptions} value={view} onChange={changeView}/>
          </div> : null
        }
        {
          showCheckbox ? <div className={styles.viewRadioMode}>
            <div>
              <span>Show Services</span>
              <Tooltip content="if the network communicate by Kubernetes service, the service name will be shown。">
                <Icon name="question-circle" />
              </Tooltip>
            </div>
            <RadioButtonGroup options={showServiceOptions} value={showService} onChange={changeShowService}/>
          </div> : null
        }
        <TopoLegend typeList={nodeTypesList} metric={lineMetric} volumes={volumes} options={options}/>
      </div>
      <div id="kindling_topo" style={{ height: '100%' }} ref={graphRef}></div>
      {
        loading ? <div className={styles.spinner_warp}>
          <Spinner className={styles.spinner_icon}/>
        </div> : null
      }
    </div>
  );
};

const getStyles = stylesFactory(() => {
  return {
    wrapper: css`
      position: relative;
    `,
    topLineMetric: css`
      position: absolute;
      top: 0;
      left: 0;
      z-index: 10;
      display: flex;
      flex-direction: column;
    `,
    metricSelect: css`
      display: flex;
      align-items: center;
      margin-bottom: 10px;
    `,
    topRightWarp: css`
      position: absolute;
      top: 0;
      right: 0;
      z-index: 10;
      display: flex;
      flex-direction: column;
      width: 245px;
    `,
    viewRadioMode: css`
      display: flex;
      align-items: center;
      justify-content: space-between;
      margin-bottom: 10px;
    `,
    svg: css`
      position: absolute;
      top: 0;
      left: 0;
    `,
    textBox: css`
      position: absolute;
      bottom: 0;
      left: 0;
      padding: 10px;
    `,
    spinner_warp: css`
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background-color: #181b1fc2;
      z-index: 20;
    `,
    spinner_icon: css`
      position: absolute;
      font-size: xx-large;
      top: 48%;
      left: 49%;
    `,
  };
});
