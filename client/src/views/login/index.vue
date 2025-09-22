<template>
  <div class="select-none">
    <img :src="bg" class="wave" />
    <div class="flex-c absolute right-5 top-3">
      <!-- 主题 -->
      <!-- <el-switch
        v-model="dataTheme"
        inline-prompt
        :active-icon="dayIcon"
        :inactive-icon="darkIcon"
        @change="dataThemeChange"
      /> -->
      <!-- 国际化 -->
      <!-- <el-dropdown trigger="click">
        <globalization
          class="hover:text-primary hover:bg-[transparent]! w-[20px] h-[20px] ml-1.5 cursor-pointer outline-hidden duration-300"
        />
        <template #dropdown>
          <el-dropdown-menu class="translation">
            <el-dropdown-item
              :style="getDropdownItemStyle(locale, 'zh')"
              :class="['dark:text-white!', getDropdownItemClass(locale, 'zh')]"
              @click="translationCh"
            >
              <IconifyIconOffline
                v-show="locale === 'zh'"
                class="check-zh"
                :icon="Check"
              />
              简体中文
            </el-dropdown-item>
            <el-dropdown-item
              :style="getDropdownItemStyle(locale, 'en')"
              :class="['dark:text-white!', getDropdownItemClass(locale, 'en')]"
              @click="translationEn"
            >
              <span v-show="locale === 'en'" class="check-en">
                <IconifyIconOffline :icon="Check" />
              </span>
              English
            </el-dropdown-item>
          </el-dropdown-menu>
        </template>
      </el-dropdown> -->
    </div>
    <div class="login-container">
      <div class="img">
        <component :is="toRaw(illustration)" />
      </div>
      <div class="login-box">
        <div class="login-form">
          <!-- <avatar class="avatar" /> -->
          <Motion>
            <h2 class="outline-hidden">{{ title }}</h2>
          </Motion>

          <el-form
            v-if="currentPage === 0"
            ref="ruleFormRef"
            :model="ruleForm"
            :rules="loginRules"
            size="large"
          >
            <Motion :delay="100">
              <el-form-item
                :rules="[
                  {
                    required: true,
                    message: transformI18n($t('login.pureUsernameReg')),
                    trigger: 'blur'
                  }
                ]"
                prop="username"
              >
                <el-input
                  v-model="ruleForm.username"
                  clearable
                  placeholder="邮箱"
                  :prefix-icon="useRenderIcon(User)"
                />
              </el-form-item>
            </Motion>

            <Motion :delay="150">
              <el-form-item prop="password">
                <el-input
                  v-model="ruleForm.password"
                  clearable
                  show-password
                  :placeholder="t('login.purePassword')"
                  :prefix-icon="useRenderIcon(Lock)"
                />
              </el-form-item>
            </Motion>

            <!--验证码-->
            <Motion :delay="200">
              <el-form-item prop="verifyCode">
                <el-input
                  v-model="ruleForm.verifyCode"
                  clearable
                  :placeholder="t('login.pureVerifyCode')"
                  :prefix-icon="useRenderIcon(Keyhole)"
                >
                  <template v-slot:append>
                    <ReImageVerify v-model:code="imgCode" />
                  </template>
                </el-input>
              </el-form-item>
            </Motion>

            <Motion :delay="250">
              <el-button
                class="w-full mt-4! mb-4!"
                size="default"
                type="primary"
                :loading="loading"
                :disabled="disabled"
                @click="onLogin(ruleFormRef)"
              >
                {{ t("login.pureLogin") }}
              </el-button>
            </Motion>

            <Motion :delay="300">
              <el-form-item>
                <div class="w-full h-[20px] flex justify-between items-center">
                  <el-button
                    class="w-full mt-4!"
                    size="default"
                    @click="useUserStoreHook().SET_CURRENTPAGE(2)"
                  >
                    邮箱注册
                  </el-button>

                  <!--
                  <el-button
                    class="w-full mt-4!"
                    size="default"
                    @click="useUserStoreHook().SET_CURRENTPAGE(1)"
                  >
                    手机登录
                  </el-button>
                  <el-button
                    class="w-full mt-4!"
                    size="default"
                    @click="useUserStoreHook().SET_CURRENTPAGE(5)"
                  >
                    账号注册
                  </el-button>
                  -->
                </div>
              </el-form-item>
            </Motion>

            <!--<Motion :delay="250">
              <el-form-item>
                <div class="w-full h-[20px] flex justify-center items-center">
                  <el-button
                    link
                    type="primary"
                    @click="useUserStoreHook().SET_CURRENTPAGE(4)"
                  >
                    {{ t("login.pureForget") }}
                  </el-button>
                </div>
              </el-form-item>
            </Motion>-->
          </el-form>

          <!-- 手机号注册/登录 -->
          <!--<PhoneRegister v-if="currentPage === 1" /> -->
          <!-- 邮箱注册/登录 -->
          <MailRegister v-if="currentPage === 2" />
          <!-- 忘记密码 -->
          <!--<LoginUpdate v-if="currentPage === 4" />-->
          <!-- 账号密码注册 -->
          <!--<AccountRegister v-if="currentPage === 5" />-->
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useI18n } from "vue-i18n";
import Motion from "./utils/motion";
import { useRouter } from "vue-router";
import { message } from "@/utils/message";
import { loginRules } from "./utils/rule";
import { ref, reactive, toRaw, computed } from "vue";
import { debounce } from "@pureadmin/utils";
import { useNav } from "@/layout/hooks/useNav";
import { useEventListener } from "@vueuse/core";
import type { FormInstance } from "element-plus";
import { $t, transformI18n } from "@/plugins/i18n";
import { useLayout } from "@/layout/hooks/useLayout";
import { useUserStoreHook } from "@/store/modules/user";
import { addPathMatch, getTopMenu } from "@/router/utils";
import { bg, avatar, illustration } from "./utils/static";
import { useRenderIcon } from "@/components/ReIcon/src/hooks";
import { useTranslationLang } from "@/layout/hooks/useTranslationLang";
import { useDataThemeChange } from "@/layout/hooks/useDataThemeChange";
import { ReImageVerify } from "@/components/ReImageVerify";
import PhoneRegister from "@/views/login/components/PhoneRegister.vue";
import AccountRegister from "@/views/login/components/AccountRegister.vue";
import MailRegister from "@/views/login/components/MailRegister.vue";
import LoginUpdate from "@/views/login/components/LoginUpdate.vue";

import dayIcon from "@/assets/svg/day.svg?component";
import darkIcon from "@/assets/svg/dark.svg?component";
import globalization from "@/assets/svg/globalization.svg?component";
import Keyhole from "~icons/ri/shield-keyhole-line";
import Lock from "~icons/ri/lock-fill";
import Check from "~icons/ep/check";
import User from "~icons/ri/user-3-fill";
import { accountLoginApi, nodeUserEmailLoginApi } from "@/api/managerApi";
import { usePermissionStoreHook } from "@/store/modules/permission";
import { setRoles } from "@/utils/auth";

defineOptions({
  name: "Login"
});
const imgCode = ref("");
const router = useRouter();
const loading = ref(false);
const disabled = ref(false);
const ruleFormRef = ref<FormInstance>();

const currentPage = computed(() => {
  return useUserStoreHook().currentPage;
});
const { initStorage } = useLayout();
initStorage();

const { t } = useI18n();
const { dataTheme, overallStyle, dataThemeChange } = useDataThemeChange();
dataThemeChange(overallStyle.value);
const { title, getDropdownItemStyle, getDropdownItemClass } = useNav();
const { locale, translationCh, translationEn } = useTranslationLang();

const ruleForm = reactive({
  // username: "admin",
  // password: "admin123",
  username: "",
  password: "",
  verifyCode: ""
});

const onLogin = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(valid => {
    if (valid) {
      loading.value = true;
      // useUserStoreHook()
      //   .loginByUsername({
      //     username: ruleForm.username,
      //     password: ruleForm.password
      //   })
      const tempTime = useUserStoreHook()?.imageCodeTime || "";
      nodeUserEmailLoginApi({
        mail: ruleForm.username,
        password: ruleForm.password,
        graph_code: ruleForm.verifyCode,
        t: tempTime
      })
        .then(res => {
          console.log("res ========= ", res);
          console.log("admin -==== ", res.data.user_info.user_name);
          if (res.errcode === 0) {
            if (res.data.user_info.user_name) {
              setRoles(
                res.data.user_info.is_admin,
                res.data.user_info.user_name
              );
            } else {
              setRoles(
                res.data.user_info.is_admin,
                "用户" + String(res.data.user_info.id)
              );
            }

            usePermissionStoreHook().handleWholeMenus([]);
            addPathMatch();
            console.log(getTopMenu(true).path, "          new path");
            // router.push(getTopMenu(true).path);
            router.push("/welcome");
            message("登录成功", { type: "success" });
            loading.value = false;
          }
        })
        .finally(() => (loading.value = false));

      // accountLoginApi({
      //   user_name: ruleForm.username,
      //   password: ruleForm.password,
      //   graph_verify_code: ruleForm.verifyCode,
      //   t: tempTime
      // })
      //   .then(res => {
      //     console.log("res ========= ", res);
      //     console.log("admin -==== ", res.data.user_info.user_name);
      //     if (res.errcode === 0) {
      //       setRoles(res.data.user_info.is_admin, res.data.user_info.user_name);
      //       usePermissionStoreHook().handleWholeMenus([]);
      //       addPathMatch();
      //       console.log(getTopMenu(true).path, "          new path");
      //       // router.push(getTopMenu(true).path);
      //       router.push("/welcome");
      //       message("登录成功", { type: "success" });
      //       loading.value = false;
      //     }
      //   })
      //   .finally(() => (loading.value = false));
    }
  });
};

const immediateDebounce: any = debounce(
  formRef => onLogin(formRef),
  1000,
  true
);

useEventListener(document, "keydown", ({ code }) => {
  if (
    ["Enter", "NumpadEnter"].includes(code) &&
    !disabled.value &&
    !loading.value
  )
    immediateDebounce(ruleFormRef.value);
});
</script>

<style scoped>
@import url("@/style/login.css");
</style>

<style lang="scss" scoped>
:deep(.el-input-group__append, .el-input-group__prepend) {
  padding: 0;
}

.translation {
  ::v-deep(.el-dropdown-menu__item) {
    padding: 5px 40px;
  }

  .check-zh {
    position: absolute;
    left: 20px;
  }

  .check-en {
    position: absolute;
    left: 20px;
  }
}
</style>
