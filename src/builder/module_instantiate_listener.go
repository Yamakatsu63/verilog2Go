package builder

import (
	parser "github.com/verilog2Go/antlr/verilog"
)

type Instance struct {
	moduleName   string
	instanceName string
	//ordered_portの場合はports、Named_portの場合はportMapを使う
	ports   []string
	portMap map[string]string
}

var instance Instance

// EnterModule_instantiation is called when production module_instantiation is exited.
func (s *CustomVerilogListener) EnterModule_instantiation(ctx *parser.Module_instantiationContext) {
	//インスタンスの初期化
	instance = Instance{}
}

// ExitModule_instantiation is called when production module_instantiation is exited.
func (s *CustomVerilogListener) ExitModule_instantiation(ctx *parser.Module_instantiationContext) {
	//インスタンスの情報をBuilderに渡す
	instance.moduleName = ctx.Module_identifier().GetText()
	CreateInstance(instance)
}

// ExitModule_instance is called when production module_instance is exited.
func (s *CustomVerilogListener) ExitModule_instance(ctx *parser.Module_instanceContext) {
	instance.instanceName = ctx.Name_of_instance().GetText()
}

// ExitOrdered_port_connection is called when production ordered_port_connection is exited.
func (s *CustomVerilogListener) ExitOrdered_port_connection(ctx *parser.Ordered_port_connectionContext) {
	instance.ports = append(instance.ports, ctx.Expression().GetText())
}

// ExitNamed_port_connection is called when production named_port_connection is entered.
func (s *CustomVerilogListener) ExitNamed_port_connection(ctx *parser.Named_port_connectionContext) {
}
