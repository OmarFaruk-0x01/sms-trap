// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package guide

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"OmarFaruk-0x01/sms-trap/app/views/layouts"
)

type ShowGuidesProps struct {
	InboxLayoutProps *layouts.InboxLayoutProps
}

func ShowGuides(props *ShowGuidesProps) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Var2 := templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
			if !templ_7745c5c3_IsBuffer {
				templ_7745c5c3_Buffer = templ.GetBuffer()
				defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"relative h-full overflow-y-auto\"><div class=\"sticky top-0 bg-white z-50\"><div class=\"px-6 py-4\"><h2 class=\"text-2xl font-bold text-gray-800\">Framework Guides</h2><p class=\"text-sm `text-gray-600 mt-1\">Integrate SMS Trap in popular frameworks</p></div><hr class=\"border-gray-200\"><div class=\"grid grid-cols-3 gap-3 p-5\"><a href=\"https://github.com/OmarFaruk-0x01/sms-trap/tree/master/examples/laravel#sms-trap-laravel-integration\" target=\"_blank\" class=\"p-5 rounded bg-white border border-gray-300 hover:border-primary-500 hover:scale-105 duration-300 transition-all flex items-center justify-center gap-3 flex-col\"><img width=\"50px\" src=\"/static/laravel.svg\"> <span class=\"font-semibold text-lg\">Laravel</span></a> <a href=\"https://github.com/OmarFaruk-0x01/sms-trap/tree/master/examples/django#sms-trap-django-integration\" target=\"_blank\" class=\"p-5 rounded bg-white border border-gray-300 hover:border-primary-500 hover:scale-105 duration-300 transition-all flex items-center justify-center gap-3 flex-col\"><img width=\"100px\" class=\"w-[100px]\" src=\"/static/django.svg\"> <span class=\"font-semibold text-lg\">Django</span></a> <a class=\"p-5 rounded bg-white border border-gray-300 hover:border-primary-500 hover:scale-105 duration-300 transition-all flex items-center justify-center gap-3 flex-col\"><img width=\"100px\" class=\"w-[100px]\" src=\"/static/rails.svg\"> <span class=\"font-semibold text-lg\">Ruby On Rails</span></a></div></div><div class=\"px-6 py-8 space-y-8\"></div></section>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			if !templ_7745c5c3_IsBuffer {
				_, templ_7745c5c3_Err = io.Copy(templ_7745c5c3_W, templ_7745c5c3_Buffer)
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = layouts.InboxLayout(props.InboxLayoutProps).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

// templ apiDetails(d *GuidePageProps) {
// 	<div class="bg-white shadow rounded-lg overflow-hidden">
// 		<div class="px-6 py-4 border-b border-gray-200">
// 			<h3 class="text-lg font-semibold text-gray-800">API Details</h3>
// 		</div>
// 		<div class="px-6 py-4">
// 			<table class="w-full text-sm text-left text-gray-500">
// 				<tbody>
// 					<tr class="border-b">
// 						<th class="py-3 font-medium text-gray-900">Host</th>
// 						<td class="py-3">localhost:{ d.Port }</td>
// 					</tr>
// 					<tr class="border-b">
// 						<th class="py-3 font-medium text-gray-900">Base URL</th>
// 						<td class="py-3">http://localhost:{ d.Port }/api/v1</td>
// 					</tr>
// 				</tbody>
// 			</table>
// 		</div>
// 	</div>
// }

// templ apiEndpoints() {
// 	<div class="bg-white shadow rounded-lg overflow-hidden">
// 		<div class="px-6 py-4 border-b border-gray-200">
// 			<h3 class="text-lg font-semibold text-gray-800">API Endpoints</h3>
// 		</div>
// 		<div class="px-6 py-4 flex flex-col space-y-6">
// 			@endpointCard("GET", "/trap", "Trap SMS", []endpointParams{
// 				endpointParams{
// 					label:       "phones[]",
// 					description: "List of phones",
// 				},
// 				endpointParams{
// 					label:       "message",
// 					description: "SMS message",
// 				},
// 				endpointParams{
// 					label:       "label",
// 					description: "Type of SMS (transactional/promotional)",
// 				},
// 			})
// 		</div>
// 	</div>
// }

// type endpointParams struct {
// 	label       string
// 	description string
// }

// templ endpointCard(method string, path string, description string, params []endpointParams) {
// 	<div class="border border-gray-200 rounded-md">
// 		<div class="px-4 py-3 bg-gray-50 border-b border-gray-200 flex items-center justify-between">
// 			<span class="font-mono text-sm font-semibold">{ method } { path }</span>
// 			<span class="px-2 py-1 text-xs font-semibold text-white bg-blue-500 rounded-full">{ method }</span>
// 		</div>
// 		<div class="px-4 py-3">
// 			<p class="text-sm text-gray-600">{ description }</p>
// 			if len(params) > 0 {
// 				<div class="mt-3">
// 					<h4 class="text-sm font-semibold text-gray-700">Queries:</h4>
// 					<ul class="mt-2 space-y-1">
// 						for _, param := range params {
// 							<li class="text-sm text-gray-600 flex items-center gap-2">
// 								<code class="px-1 py-0.5 text-sm font-mono bg-gray-100 rounded">{ param.label }</code>
// 								<span>-</span>
// 								<span>
// 									{ param.description }
// 								</span>
// 							</li>
// 						}
// 					</ul>
// 				</div>
// 			}
// 		</div>
// 	</div>
// }

// templ codeSamples(d *GuidePageProps) {
// 	<div class="bg-white shadow rounded-lg overflow-hidden">
// 		<div class="px-6 py-4 border-b border-gray-200">
// 			<h3 class="text-lg font-semibold text-gray-800">Code Samples</h3>
// 		</div>
// 		<div class="px-6 py-4 space-y-6">
// 			@core.Tabs(&core.TabsProps{
// 				TabItems: []*core.TabItemProps{
// 					&core.TabItemProps{
// 						Id:        "curl",
// 						ActiveTab: "",
// 						Label:     "cURL",
// 						Icon: func() templ.Component {
// 							return nil
// 						},
// 						Panel: func() templ.Component {
// 							return components.CodeBlock(curlExample(d.Port), "bash")
// 						},
// 					},
// 					&core.TabItemProps{
// 						Id:        "go",
// 						ActiveTab: "",
// 						Label:     "Go",
// 						Icon: func() templ.Component {
// 							return nil
// 						},
// 						Panel: func() templ.Component {
// 							return components.CodeBlock(goExample(d.Port), "go")
// 						},
// 					},
// 					&core.TabItemProps{
// 						Id:        "js",
// 						ActiveTab: "",
// 						Label:     "Javascript",
// 						Icon: func() templ.Component {
// 							return nil
// 						},
// 						Panel: func() templ.Component {
// 							return components.CodeBlock(javascriptExample(d.Port), "javascript")
// 						},
// 					},
// 					&core.TabItemProps{
// 						Id:        "python",
// 						ActiveTab: "",
// 						Label:     "Python",
// 						Icon: func() templ.Component {
// 							return nil
// 						},
// 						Panel: func() templ.Component {
// 							return components.CodeBlock(pythonFlaskExample(d.Port), "python")
// 						},
// 					},
// 					&core.TabItemProps{
// 						Id:        "ruby",
// 						ActiveTab: "",
// 						Label:     "Ruby",
// 						Icon: func() templ.Component {
// 							return nil
// 						},
// 						Panel: func() templ.Component {
// 							return components.CodeBlock(rubyExample(d.Port), "ruby")
// 						},
// 					},
// 					&core.TabItemProps{
// 						Id:        "php",
// 						ActiveTab: "",
// 						Label:     "PHP",
// 						Icon: func() templ.Component {
// 							return nil
// 						},
// 						Panel: func() templ.Component {
// 							return components.CodeBlock(phpExample(d.Port), "php")
// 						},
// 					},
// 				},
// 			})
// 		</div>
// 	</div>
// }
