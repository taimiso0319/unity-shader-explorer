package shdrexpl

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/taimiso0319/unity-shader-explorer/analyze"
	"github.com/taimiso0319/unity-shader-explorer/collect"
)

const channelName = "shader_explorer"

// UnityShaderExplorerPlugin implements flutter.Plugin and handles method.
type UnityShaderExplorerPlugin struct{}

var _ flutter.Plugin = &UnityShaderExplorerPlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *UnityShaderExplorerPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("getPlatformVersion", p.handlePlatformVersion)
	channel.HandleFunc("getShaders", p.handleCollectShader)
	return nil
}

func (p *UnityShaderExplorerPlugin) handlePlatformVersion(arguments interface{}) (reply interface{}, err error) {
	return "go-flutter " + flutter.PlatformVersion, nil
}

func (p *UnityShaderExplorerPlugin) handleCollectShader(path string) (reply interface{}, err error) {
	var shaderPaths []string = collect.GetShaderPaths(path)
	return analyze.GetShaderDetails(shaderPaths)
}
