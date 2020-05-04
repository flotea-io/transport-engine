<template>
	<el-dialog :title="'Plan wyjazdu #'+(data.id + 1)" :visible.sync="data.dialogVisible" :close-on-click-modal="false" :close-on-press-escape="false" :show-close="false" width="950px">

		<wd-repeat :data.sync="data.checkedWd" :open.sync="data.dialogVisible"/>
		<br>
		<!--<month-day-range :data.sync="data.mdRange" :open.sync="data.dialogVisible" />-->
		<br>
		<ev-range :data.sync="data.evRange" :not="false" :open.sync="data.dialogVisible"/>
		<br>
		<hr-range :data.sync="data.hrRange" :open.sync="data.dialogVisible"/>

		<ev-range :data.sync="data.notEvRange" :not="true" :open.sync="data.dialogVisible" />

		<span slot="footer" class="dialog-footer">
			<el-button :disabled="newEvent" @click="$listeners.delete" type="danger">Usuń</el-button>
			<el-button @click="$listeners.close(); data.dialogVisible = false;">Anuluj</el-button>
			<el-button @click="$listeners.save" type="success">Zapisz</el-button>
		</span>
	</el-dialog>
</template>

<script>
	import EvRange from '~/components/EvRange.vue'
	import HrRange from '~/components/HrRange.vue'
	import WdRepeat from '~/components/WdRepeat.vue'
	//import MonthDayRange from '~/components/MonthDayRange.vue'

	export default {
		components: {
			EvRange, HrRange, WdRepeat//, MonthDayRange
		},
		props: ["data", "newEvent"],
		event: ['change', 'delete'],
	}

</script>